package basic

import "fmt"

func StudyFmt() {

	var num int64
	num = 10
	str := "안녕하세요"

	// 출력

	// fmt.Print() : 입력값들을 출력한다.
	fmt.Print("# 예시")
	fmt.Print(num)
	fmt.Print(str)
	fmt.Print("\n")
	//결과 :
	//# 예시10안녕하세요

	//fmt.Println() : 출력하고 엔터기능까지
	fmt.Println("# 예시")
	fmt.Println(num)
	fmt.Println(str)
	//결과:
	//# 예시
	//10
	//안녕하세요

	//fmt.Printf() :포맷에 맞게 입력값을 출력한다.
	fmt.Printf("#예시\n%d\n%s\n", num, str)
	//결과:
	//# 예시
	//10
	//안녕하세요

	//입력
	//fmt.Scan() : 입력받기
	var in int
	number, err := fmt.Scan(&in)
	if err != nil {
		fmt.Println("에러 : ", err)
	}
	fmt.Println("입력 개수 : ", number)
	fmt.Println("입력값 : ", in)

	//fmt.Scanln() : 입력 받고 공란으로 구분된 값들을 읽는다.
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("에러 : ", err)
	}
	fmt.Println("입력 개수 : ", n)
	fmt.Println("입력값 : ", a, ",", b)

	//fmt.Scanf() : 포맷에 맞게 입력한다.
	var c int
	var d string
	n, err = fmt.Scanf("%d/%s", &c, &d)
	fmt.Println("입력 개수 : ", n)
	fmt.Println("입력값 : ", c, ",", d)
	if err != nil {
		fmt.Println("에러 : ", err)
	}
}
