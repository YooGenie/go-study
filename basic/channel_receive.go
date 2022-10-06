package basic

import "fmt"

func ReceiveDataFromChannel(ch <-chan int) { //입력
	for {
		select {
		case receive, ok := <-ch:
			fmt.Println(receive)
			if !ok {
				return
			}
		}
	}
	//값을 send 할 수 없다.
}
