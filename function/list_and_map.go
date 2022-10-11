package function

import (
	"fmt"
)

type result struct {
	Id               int64
	Name             string
	RegisterOrgCount int64
	ReportCount      int64
}

func ListAndMap() {
	var list []result

	list = append(list, result{
		Id:   1,
		Name: "유지니",
	})

	list = append(list, result{
		Id:   2,
		Name: "박재범",
	})

	list = append(list, result{
		Id:   3,
		Name: "이수혁",
	})

	//resultList *[]result 타입으로 만들었다.
	// 구조체에 포인터를 준 이유?
	// 전체 구조체를 복사하는 것을 피하기 위해 포인터에 작용했다.
	var resultList *[]result
	resultList = &list
	fmt.Println("대입 전 resultList: ", resultList)

	addValue3 := map[string]int64{
		"id":               1,
		"registerOrgCount": 4,
		"reportCount":      1,
	}

	addValue4 := map[string]int64{
		"id":               2,
		"registerOrgCount": 2,
		"reportCount":      2,
	}

	//맵이 배열로 포인트로 있다
	//registerOrgCount : &[map[id:2 registerOrgCount:2 reportCount:2] map[id:1 registerOrgCount:4 reportCount:1]]
	var registerOrgCount []map[string]int64
	registerOrgCount = append(registerOrgCount, addValue4)
	registerOrgCount = append(registerOrgCount, addValue3)
	resultCount := &registerOrgCount
	fmt.Println(resultCount)

	//resultList 안에 id에 맞게 resultCount 대입 시킨다.
	for i := range *resultList {
		for _, v := range *resultCount {
			if (*resultList)[i].Id == v["id"] {
				(*resultList)[i].RegisterOrgCount = v["registerOrgCount"]
			}
		}
		// resultList[i].Id =1 에러 발생
		//Invalid operation: 'resultList[i]' (type '*[]result' does not support indexing)
		//resultList 포인터형이라서 값을 가져오기 위해서 (*resultList) 해야한다.
	}
	fmt.Println("대입 후 resultList: ", resultList)
}
