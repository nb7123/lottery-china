package model

import (
	"time"
	"sort"
)

const (
	LotteryTypeWelfare = 1
	LotteryTypeSSC = 2
)

type Lottery struct {
	Num int
	OpeningTime time.Time
	ResultNum []int
	Type int
}

type CountResult struct {
	Count int
	Numbers []int
	RLottery []Lottery
}

type CountResultSorter struct {
	data []CountResult
	by func(p1, p2 *CountResult) bool
}

func (sorter *CountResultSorter) Len() int {
	return len(sorter.data)
}

func (sorter *CountResultSorter) Swap(i, j int) {
	sorter.data[i], sorter.data[j] = sorter.data[j], sorter.data[i]
}

func (sorter *CountResultSorter) Less(i, j int) bool {
	return sorter.by(&sorter.data[i], &sorter.data[j])
}

type By func(p1, p2 *CountResult) bool

func (by By) Sort(data []CountResult) {
	sorter := &CountResultSorter{
		data: data,
		by: by,
	}

	sort.Sort(sorter)
}
