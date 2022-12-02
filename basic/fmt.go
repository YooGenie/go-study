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

func StudyFmtError() {
	var in int
	_, err := fmt.Scan(&in)
	if err != nil {
		fmt.Println("에러 : ", err)
	}
	/*
		에러 메시지
		1) expected integer : int 인데 string 입력했을 때
	*/

	var a int
	var b int

	_, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("에러 : ", err)
	}

	/*
		에러 메시지
		1) unexpected newline : 2개 입력해야하는데 1개만 입력하고 엔터 쳤을 때
		예시로 7만 입력하고 엔터 치면 에러

	*/

	var c int
	var d string
	_, err = fmt.Scanf("%d/%s", &c, &d)
	if err != nil {
		fmt.Println("에러 : ", err)
	}
	/*
		에러 메시지
		1) input does not match format  : 포맷 맞지 않을 때 1,2을 입력해야했는데 1 2 입력하면 에러
	*/

}
