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

	Solution()
	BlockCaseFull()
	BlockCaseEmpty()
}

func BlockCaseFull() {
	/*
		1,2,3 넣으면 채널이 꽉 찼다. 이 상태에서 4를 넣으면 어떻게 될까?
		=> 채널이 꽉 차서 4는 send하기 위해서 계속 대기를 탄다.
		=> finish를 출력하지 못하고 프로그램이 비정상적으로 종료 된다
		=> 아래와 같은 메시지가 뜬다
		=> fatal error: all goroutines are asleep - deadlock!
	*/
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	fmt.Println("finish")
}

func BlockCaseEmpty() {
	/*
		1,2,3를 넣고 receive를 4번 하면 어떻게 될까?
		=> 4번째에서 채널안에 값이 없어서 가져올 수 없는 상황이다.
		=> 그래서 무한 대기를 타다가 비정상적으로 종료가 된다
		=> 아래와 같은 메시지가 뜬다
		=> fatal error: all goroutines are asleep - deadlock!
	*/
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	<-ch         // 값을 꺼내지만 사용하지 않는 경우 이다.
	num1 := <-ch //값을 꺼내서 변수에 넣어주는 경우
	fmt.Println(num1)
	<-ch
	<-ch //채널에 값을 가져올 수 없어서 무한 대기

	fmt.Println("finish")
}

func Solution() {
	// close를 해주면 채널에 값이 없는 경우 zero 값을 가져온다.
	// 채널의 데이터타입이 string이면 "" 빈 문자열을 가져온다.
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) //값이 빠져 나가면 0으로 채워준다.

	<-ch
	<-ch
	<-ch
	result := <-ch //0이 빼와서 에러가 나지 않고 종료 된다
	fmt.Println("result :", result)
	fmt.Println("finish")

	ch <- 5 // close하면 send 할 수 없다. panic: send on closed channel
}
