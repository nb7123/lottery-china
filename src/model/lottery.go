package model

import "time"

const (
	LotteryTypeWelfare = 1
	LotteryTypeSSC = 2
)

type Lottery struct {
	Num string
	OpeningTime time.Time
	ResultNum []int
	Type int
}

