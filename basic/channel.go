package basic

import "fmt"

func Channel() {
	/*
		채널은 값을 전달하는 통로이다.
		고루틴과 고루틴 사이에 값을 전달해준다.
	*/

	ch := make(chan int, 3) // 채널 생성 방법 (데이트타입, size) ex) int 타입의 채널에 데이터를 3개 넣을 수 있다.

	//채널 ch에 값을 넣기
	ch <- 1
	ch <- 2

	//채널 ch에서 값을 빼기
	received1 := <-ch
	received2 := <-ch

	fmt.Println(received1)
	fmt.Println(received2)

}