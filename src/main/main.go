package main

import (
	"fmt"
	"time"
	"dbhelper"
	"model"
	"log"
	"os"
)

const (
	startId = 18001
)

func main() {
	fmt.Println("Hello Lottery!")

	dbhelper.InitDB()

	//err := lottery.UpdateHistory()
	//
	//if nil != err {
	//	log.Fatal(err)
	//}

	//countWelfare1()
	//countWelfare2()
	result, err := countWelfare3()
	if nil == err {
		model.By(func(p1, p2 *model.CountResult) bool {
			return p1.Count > p2.Count
		}).Sort(result)
	}

	if nil == err {
		f, err := os.OpenFile("./3.log", os.O_CREATE|os.O_RDWR, 0766)
		if nil != err {
			log.Fatal(err)
		}

		logger := log.New(f, "", log.LstdFlags)

		//logger.Printf("Used time[%v]\nResult is: %v", time.Now().UnixNano() - startTime, result)
		logger.Printf("从%v期开始到现在数据统计", startId)
		for _, item := range result {
			var ids []int
			for _, l := range item.RLottery {
				ids = append(ids, l.Num)
			}

			logger.Printf("出现次数[%2d], 数字序列%2v, 关联期号%5v", item.Count, item.Numbers, ids)
		}
	} else {
		log.Fatal(err)
	}

	log.Printf("Count[%v]", len(result))
	//getSSCData()
}

func getSSCData() {
	var err error
	var lastID = 0
	start := time.Now()

	// 还没有数据，开始时间为2010年11月11号
	if nil != err {
		fmt.Println(err)
		start = time.Date(2010, time.November, 11, 0, 0, 0, 0, time.Local)

	} else {
		// 计算日期
		// 1、取掉最后三位期号
		// 2、得到最后两位天
		// 3、得到最后两位月
		// 4、得到年份
		date := lastID / 1000

		day := date % 100
		date /= 100

		month := date % 100
		date /= 100

		year := date

		fmt.Printf("LastID[%d], Year[%d], Month[%d], Day[%d]", lastID, year, month, day)

		start = time.Date(year + 2000, time.Month(month), day, 0, 0, 0, 0, time.Local)
	}

	fmt.Printf("Start[%v]\n", start)
	//ssc.GetAllHistoryData(start)
}

func countWelfare1() ([]model.CountResult, error) {
	var result []model.CountResult
	var err error = nil

	var tmp []model.Lottery
	for i := 1; i <= 33; i++ {
		tmp, err = dbhelper.SearchNumberCount1(i, startId)
		if nil != err {
			break
		}

		result = append(result, model.CountResult{Count: len(tmp), Numbers: []int{i}, RLottery: tmp})
	}

	return result, err
}

func countWelfare2() ([]model.CountResult, error) {
	var result []model.CountResult
	var err error = nil
	var tmp []model.Lottery
	for i := 1; i <= 33; i++ {
		for j := 1; j < i; j++  {
			if i == j {
				continue
			}
			tmp, err = dbhelper.SearchNumberCount2(i, j, startId)
			if nil != err {
				break
			}

			result = append(result, model.CountResult{Count: len(tmp), Numbers: []int{i, j}, RLottery: tmp})
		}
	}

	return result, err
}

func countWelfare3() ([]model.CountResult, error) {
	var result []model.CountResult
	var err error = nil

	var tmp []model.Lottery
	for i := 1; i <= 33; i++ {
		for j := 1; j <= i; j++  {
			if i == j {
				continue
			}
			for k := 1; k < j; k++ {
				if k == j {
					continue
				}

				tmp, err = dbhelper.SearchNumberCount3(i, j, k, startId)
				if nil != err {
					break
				}


				result = append(result, model.CountResult{Count: len(tmp), Numbers: []int{i, j, k}, RLottery: tmp})
			}
		}
	}

	return result, err
}

func countWelfare4() ([]model.CountResult, error) {
	var result []model.CountResult
	var err error = nil

	var tmp []model.Lottery
	for i := 1; i <= 33; i++ {
		for j := 1; j <= i; j++  {
			if i == j {
				continue
			}
			for k := 1; k < j; k++ {
				if k == j {
					continue
				}

				for l := 1; l < k; l++ {
					tmp, err = dbhelper.SearchNumberCount4(i, j, k, l, startId)
					if nil != err {
						break
					}


					result = append(result, model.CountResult{Count: len(tmp), Numbers: []int{i, j, k, l}, RLottery: tmp})
				}
			}
		}
	}

	return result, err
}
