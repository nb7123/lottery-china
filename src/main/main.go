package main

import (
	"fmt"
	"time"
	"dbhelper"
	"lottery"
)

func main() {
	fmt.Println("Hello Lottery!")

	dbhelper.InitDB()

	lottery.UpdateHistory()

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
