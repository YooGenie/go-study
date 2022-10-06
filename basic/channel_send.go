package basic

import "time"

func SendDataToChannel() <-chan int { //값을 넣을 수 있다
	ch := make(chan int, 3)
	go func() {
		time.Sleep(time.Second)
		ch <- 3
	}()

	return ch //receive 할 수 있는 채널로 리턴 해준다.
}
