package basic

import "fmt"

func Goroutine() {
	fmt.Println("Goroutine 시작")
	go func() {
		fmt.Println("Goroutine 진행 중")
	}()
	go add()
	fmt.Println("Goroutine 끝")

}

func add() {
	sum := 1 + 2
	fmt.Printf("합계 : %d", sum)

}
