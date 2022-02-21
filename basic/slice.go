package basic

import "fmt"

func SliceStudy() {
	var slice []int // 배열과 다르게 크기를 지정하지 않는다.
	// 길이가 0이 배열이 생성 된다.

	// 길이 체크
	if len(slice) == 0 {
		fmt.Println("현재 슬라이스 길이는 0 입니다.")
	}

	//slice[0] =1
	//fmt.Println("슬라이스 0 : ",slice[0]) //패닉 발생
	// 길이가 0인데 첫번째 칸에 값을 넣으려고 하니까 에러가 발생 된 것이다.

	// 초기화 만드는 법
	//1) {}를 이용하는 방법
	nums1 := []int{1, 2, 3}
	nums2 := []int{1: 5, 10: 2, 15: 3}
	fmt.Println("nums1 결과값: ", nums1)
	fmt.Println("nums2 결과값: ", nums2)

	str := []string{"안녕하세요", "유지니", 5: "하하하"}
	boolean := []bool{true, 3: false}
	float := []float64{3.0, 5: 5.2}
	fmt.Println("string 예시 :", str)
	fmt.Println("boolean 예시 :", boolean)
	fmt.Println("float 예시 :", float)

	// 2)make 이용하는 방법
	var m = make([]int, 3) //int의 기본값은 0
	m[1] = 5
	fmt.Println("make 의 예시 : ", m)

	//슬라이스 읽기
	for i := 0; i < len(nums1); i++ {
		fmt.Println("for문 값:", nums1[i])
	}

	for _, v := range nums1 {
		fmt.Println("for range 값: ", v)
	}

	//슬라이스 추가

	// 1) 기존 변수를 복사해서 다른 변수에 넣고 값을 추가 하는 방법
	add := []int{1, 2}
	c := append(add, 6)
	fmt.Println("add 값 : ", add)
	fmt.Println("c 값 : ", c)

	// 2) 기존 변수에 바로 추가하는 방법
	add2 := []int{1, 2, 3}
	add2 = append(add2, 6)
	fmt.Println("add2 값 : ", add2)

	// 여러가지 값을 추가
	add3 := []int{1, 2}
	add3 = append(add3, 6, 4, 2)
	fmt.Println("add3 값 : ", add3)

	var add4 []int
	fmt.Println("for문 전 add4 값 : ", add4)
	for i := 0; i < 10; i++ {
		add4 = append(add4, i)
	}
	fmt.Println("for문 후 add4 값 : ", add4)

}
