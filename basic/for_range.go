package basic

import (
	"fmt"
	"reflect"
)

func StudyForRange() {
	//배열 예시
	nums := [5]int{1, 2, 3}
	for index, value := range nums {
		fmt.Println("index : ", index, " value : ", value)
	}

	//문자열 예시
	str := "안녕! Genie"
	for index, value := range str {
		fmt.Println("index : ", index, " value : ", value, "rune 타입을 string : ", string(value))
		fmt.Println(reflect.TypeOf(value))
	}

	//슬라이스 예시
	slice := []int{1, 2, 3}
	for index, value := range slice {
		fmt.Println("index : ", index, " value : ", value)
	}

	// 맵 예시
	maps := map[string]string{"name": "Genie", "age": "20"}
	for key, value := range maps {
		fmt.Println("index : ", key, " value : ", value)
	}

}
