package main

import "fmt"

func StudyIf(color string, age int) {
	if color == "pink" {
		fmt.Println("핑크입니다")
	} else if color == "red" {
		fmt.Println("빨간색입니다")
	} else {
		fmt.Println("이상한색입니다")
	}

	if age >= 20 && age < 30 {
		if age >= 20 && age <= 23 {
			fmt.Println("20대 초반입니다")
		} else if age >= 24 && age <= 27 {
			fmt.Println("20대 중반입니다")
		} else {
			fmt.Println("20대 후반입니다")
		}
	} else if age >= 30 && age < 40 {
		if age >= 30 && age <= 33 {
			fmt.Println("30대 초반입니다")
		} else if age >= 34 && age <= 37 {
			fmt.Println("30대 중반입니다")
		} else {
			fmt.Println("30대 후반입니다")
		}
	} else if age < 20 {
		fmt.Println("10대입니다")
	} else {
		fmt.Println("40대이상입니다")
	}

	if totalCount, err := total(5); err != "" {
		fmt.Println("에러가 발생했습니다")
	} else if totalCount == 0 {
		fmt.Println("토탈값이 없습니다")
	} else {
		fmt.Println("토탈값 : ", totalCount)
	}
}

func total(num int) (totalCount int, err string) {
	if num > 0 && num < 10 {
		totalCount = num * num
		return
	} else {
		err = "에러입니다"
		return
	}
}
