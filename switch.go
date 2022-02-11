package main

import "fmt"

func StudySwitch(today string, day int64) {

	// 기본 switch 문 구조
	switch today {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("오늘은 출근해야합니다")
	case "saturday", "sunday":
		fmt.Println("오늘은 푹 쉬어도 됩니다")
	default:
		fmt.Println("값을 잘못입력했습니다.")
	}

	//비굣값없는 switch 문
	switch {
	case day >= 1 && day <= 10:
		fmt.Println("월 초 입니다")
	case day > 20 && day <= 31:
		fmt.Println("월 말입니다.")
	case day > 31:
		fmt.Println("잘못입력했습니다.")
	default:
		fmt.Println("월 중간입니다.")
	}

	//초기값 설정하는 switch 문
	switch month := 10; month {
	case 1, 2, 12:
		fmt.Println("겨울입니다.")
	case 3, 4, 5:
		fmt.Println("봄입니다.")
	case 6, 7, 8:
		fmt.Println("여름입니다.")
	case 9, 10, 11:
		fmt.Println("가을입니다.")
	default:
		fmt.Println("1~12월까지만 존재합니다.")
	}

}
