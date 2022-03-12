package basic

import "fmt"

func Map() {
	// 기본 문법

	//map[키타입]값타입{}
	ex1 := map[string]int64{}
	fmt.Println("map[string]int64{} : ", ex1)
	//map[키타입]값타입{초기값}
	ex2 := map[string]int64{"id": 2, "count": 1}
	fmt.Println("map[키타입]값타입{초기값} : ", ex2)
	//make(map[키타입]값타입, 초기용량)
	ex3 := make(map[string]int64, 0)
	fmt.Println("make(map[string]int64, 0) : ", ex3)
	//make(map[키타입]값타입)
	ex4 := make(map[string]int64)
	fmt.Println("make(map[string]int64) : ", ex4)

	//값추가
	addValue := map[string]int64{}
	//1
	addValue["id"] = 5
	addValue["count"] = 4
	fmt.Println("첫번째 추가 방법 : ", addValue)
	addValue2 := map[string]int64{"id": 4, "count": 1}
	fmt.Println("두번째 추가 방법 : ", addValue2)
	addValue3 := map[string]int64{
		"id":    4,
		"count": 1, //여러줄로 쓸때에는 마지막에 ,(콤마)를 붙여준다.
	}
	fmt.Println("세번째 추가 방법 : ", addValue3)

	//for range 이용해서 값 불러오기
	for index, value := range addValue3 {
		fmt.Println("인덱스 값 : ", index, "값 : ", value)
	}

	//값 찾기
	fmt.Println("addValue3[id] 값 :", addValue3["id"])

	// 요소 추가, 수정, 삭제
	fmt.Println("수정 전 : ", addValue3)
	addValue3["id"] = 9 //수정
	fmt.Println("id가 있어서 수정 후 : ", addValue3)

	fmt.Println("추가 전 :", addValue3)
	addValue3["num"] = 7
	fmt.Println("num이 없어서 추가 후 : ", addValue3)

	delete(addValue3, "id")
	fmt.Println("삭제 후 : ", addValue3)

}
