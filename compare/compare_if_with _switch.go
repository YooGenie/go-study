package compare

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/util"
	"time"
)

func CompareIfWithSwitch(month int64) {

	var season string
	startSwitch := util.CurrentTimeMillis()
	fmt.Println("Switch문 시작 : ", startSwitch)
	for ; month < 100000000; month++ {
		switch month {
		case 1, 2, 12:
			season = "겨울"
		case 3, 4, 5:
			season = "봄"
		case 6, 7, 8:
			season = "여름"
		case 9, 10, 11:
			season = "가을"
		default:
			season = "계절아님"
		}
	}
	endSwitch := util.CurrentTimeMillis()
	fmt.Println("Switch문 끝 : ", endSwitch)
	fmt.Println("Switch문 결과 :", endSwitch-startSwitch)

	time.Sleep(6000)

	startIf := util.CurrentTimeMillis()
	fmt.Println("If문 시작 : ", startIf)
	for month := 0; month < 100000000; month++ {
		if month == 1 || month == 2 || month == 12 {
			season = "겨울"
		} else if month == 3 || month == 4 || month == 5 {
			season = "봄"
		} else if month == 6 || month == 7 || month == 8 {
			season = "여름"
		} else if month == 9 || month == 10 || month == 11 {
			season = "가을"
		} else {
			season = "계절아님"
		}
	}
	endIf := util.CurrentTimeMillis()
	fmt.Println("If문 끝 : ", endIf)
	fmt.Println("If문 결과 :", endIf-startIf)

	fmt.Println(season)
}
