package lottery

import (
	"model"
	"time"
	"log"
	"dbhelper"

	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
)

func UpdateHistory() (error) {
	// 查看最后一条数据的ID，确定最后一次更新的时间，然后获取最后时间到现在的所有历史数据
	var startId = 0
	var err error = nil
	startId, err = dbhelper.SearchLastInertedLottery(model.LotteryTypeWelfare)

	if startId <= 0 {
		startId = 2003001
	}

	var endDate = time.Now()

	var newData []model.Lottery

	if nil == err {
		for ; ; {
			lottery, err := getLotteryData(startId)
			if nil != err {
				log.Print(err)

				continue
			}

			newData = append(newData, lottery)

			var tmp = lottery.OpeningTime
			if tmp.Year() == endDate.Year()&& tmp.Month() == endDate.Month() && tmp.Day() == endDate.Day() {
				break
			}
		}

		err = dbhelper.InsertLottery(newData)
	}


	return err
}

func getLotteryData(id int) (model.Lottery, error) {
	response, err := http.Get(fmt.Sprintf("http://www.17500.cn/ssq/details.php?issue=%v", id))

	var lottery model.Lottery

	var doc *goquery.Document
	if nil == err {
		doc, err = goquery.NewDocumentFromReader(response.Body)
	}

	if nil == err {
		var selection = doc.Find("center").
			Find("center").
			Find("table").
			Find("tbody").
			Find("tr").
			Find("td").Find("table")

		for ; ; {
			var nodeName = goquery.NodeName(selection)
			attr, exist := selection.Attr("border")
			log.Printf("Index: %v, Node name: %v, attribute: %v exist: %v", 1, nodeName, attr, exist)

			selection = selection.Next()
		}
		selection.Children().EachWithBreak(func(i int, s *goquery.Selection) bool {
				var nodeName = goquery.NodeName(s)
				attr, exist := s.Attr("border")
				log.Printf("Index: %v, Node name: %v, attribute: %v exist: %v", i, nodeName, attr, exist)

				s.Next()

				return true
			})
		//selection.Each()
	}

	return lottery, err
}
