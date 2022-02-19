package function

import (
	"fmt"
	"reflect"
)

func CheckType() {
	num := 3
	boolean := true
	float := 5.2222
	arr := [6]string{"월", "화", "수", "목", "금"}
	str := "안녕! Genie"

	//rune 타입이 int32로 나왔다.
	for index, value := range str {
		fmt.Println("index : ", index, " value : ", value)
		fmt.Println(reflect.TypeOf(value))
	}

	fmt.Println("string 타입 : ", reflect.TypeOf(str))
	fmt.Println("int 타입 : ", reflect.TypeOf(num))
	fmt.Println("bool 타입 : ", reflect.TypeOf(boolean))
	fmt.Println("float64 타입 : ", reflect.TypeOf(float))
	fmt.Println("배열 타입 : ", reflect.TypeOf(arr))

}
