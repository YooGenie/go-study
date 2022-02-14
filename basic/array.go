package basic

import "fmt"

func StudyArray() {
	// string  경우
	var str = [6]string{"월", "화", "수", "목", "금"}
	for i := 0; i < len(str); i++ {
		fmt.Println("str[", i, "]=", str[i])
	}

	// int 경우
	nums := [5]int{1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println("nums[", i, "]=", nums[i])
	}

	//bool 경우
	boolean := [5]bool{true, true, true}
	for i := 0; i < len(boolean); i++ {
		fmt.Println("boolean[", i, "]=", boolean[i])
	}

	//지정한 index만 값을 넣기
	a := [7]string{0: "월", 6: "일"}
	for i := 0; i < len(a); i++ {
		fmt.Println("a[", i, "]=", a[i])
	}
	a[1] = "화"
	for i := 0; i < len(a); i++ {
		fmt.Println("a[", i, "]=", a[i])
	}

	//배열요소개수 생략방법
	x := [...]int{9, 8, 7}
	for i := 0; i < len(x); i++ {
		fmt.Println("x[", i, "]=", x[i])
	}

	// for 문을 이용하여 배열 읽기
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < 7; i++ {
		fmt.Println("arr[", i, "]=", arr[i])
	}
	//for i := 0; i < 8; i++ { //에러 발생 요소개수가 맞지 않아서
	//	fmt.Println("arr[", i, "]=", arr[i])
	//}

	//해결책 1 - len 이용한다
	for i := 0; i < len(arr); i++ {
		fmt.Println("err[", i, "]=", arr[i])
	}

	//해결책2 - range 이용한다
	for i, v := range arr {
		fmt.Println(i, "=", v)
	}
}
