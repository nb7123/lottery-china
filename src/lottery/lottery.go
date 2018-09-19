package lottery

import (
	"model"
	"log"
	"dbhelper"

	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"os"
)

func UpdateHistory() (error) {
	// 查看最后一条数据的ID，确定最后一次更新的时间，然后获取最后时间到现在的所有历史数据
	var startId = 0
	var err error = nil
	startId, err = dbhelper.SearchLastInsertedWelfare()

	if startId <= 0 {
		err = initWelfare()
		startId = 18108
	}

	var currentId = 18109

	if nil == err {

		for i := startId; i < currentId; i += 100 {
			var endId = i + 100
			if endId > currentId {
				endId = currentId
			}
			lottery, err := getLotteryData(i, i+100)
			if nil != err {
				log.Fatal(err)
			}

			err = dbhelper.InsertWelfare(lottery)
			if nil != err {
				log.Fatal(err)
			}
		}
	}


	return err
}

func initWelfare() error {
	f, err := os.Open("./03001-18108.html")
	defer f.Close()

	var doc *goquery.Document
	if nil == err {
		doc, err = goquery.NewDocumentFromReader(f)
	}

	if nil == err {
		err = dbhelper.InsertWelfare(parseWelfareData(doc))
	}

	return err
}

func getLotteryData(id int, cur int) ([]model.Lottery, error) {
	var path = fmt.Sprintf("http://datachart.500.com/ssq/?expect=all&from=%05d&to=%05d", id, cur)
	log.Printf("Data path: %v", path)
	response, err := http.Get(path)

	var lottery []model.Lottery

	var doc *goquery.Document
	if nil == err {
		doc, err = goquery.NewDocumentFromReader(response.Body)
	}

	if nil == err {
		lottery = parseWelfareData(doc)
		//selection.Each()
	}

	return lottery, err
}

func parseWelfareData(doc *goquery.Document) ([]model.Lottery) {
	var lottery []model.Lottery

	var selection = doc.Find("#tdata").Children()

	selection.Each(func(i int, s *goquery.Selection) {
		log.Printf("Handle index: %v", i)
		if !s.HasClass("tdbck") {
			var tmp = model.Lottery{Type: model.LotteryTypeWelfare}
			s.Children().Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					// number
					v, err := strconv.Atoi(strings.Trim(s.Text(), " "))
					if nil != err {
						log.Fatal(s.Text(), err)
					}
					tmp.Num = v
				}

				if s.HasClass("chartBall01") || s.HasClass("chartBall02"){
					// number
					v, err := strconv.Atoi(strings.Trim(s.Text(), " "))
					if nil != err {
						log.Fatal(err)
					}

					tmp.ResultNum = append(tmp.ResultNum, v)
				}
			})

			lottery = append(lottery, tmp)

			b, _ := json.Marshal(tmp)
			log.Printf("Lottery info[%v]", string(b))
		}
	})

	return lottery
}
