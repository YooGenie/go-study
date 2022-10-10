package basic

import "fmt"

func Goroutine() {
	fmt.Println("Goroutine 시작")
	go func() {
		fmt.Println("Goroutine 진행 중")
	}()
	fmt.Println("Goroutine 끝")

}

func Add() {
	sum := 1 + 2
	fmt.Printf("합계 : %d \n", sum)
}
