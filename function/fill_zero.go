package function

import "fmt"

func FillZero(num int)  {
	fmt.Println("num의 값 : ",num)
	fmt.Printf("10자리 수 만들기 : %010d\n", num)
	fmt.Printf("8자리 수 만들기 : %08d\n", num)

	result := fmt.Sprintf("%010d", num)
	fmt.Println("result 값 : ", result)
}
