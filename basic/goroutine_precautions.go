package basic

import (
	"fmt"
	"sync"
)

// 주의사항 1) wg.Add(1) 위치가 중요하다

func BadExWaitGroup() {
	fmt.Println("나쁜 예시 시작")
	size := 10
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		go func() {
			wg.Add(1)
			fmt.Printf("진행 중 %d \n", i)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("나쁜 예시 끝")
}

func GoodExWaitGroup() {

	fmt.Println("좋은 예시 시작")
	size := 10
	var wg sync.WaitGroup
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			fmt.Printf("진행 중 %d \n", i)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("좋은 예시 끝")
}

// 주의사항 2) 변수를 조심히 해야한다.

func SolutionOne() {
	fmt.Println("해결책 one 시작")
	size := 10
	var wg sync.WaitGroup
	wg.Add(size)
	for i := 0; i < size; i++ {
		i := i
		go func() {
			fmt.Printf("진행 중 %d \n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("해결책 one  끝")
}

func SolutionTwo() {
	fmt.Println("해결책 two 시작")
	size := 10
	var wg sync.WaitGroup
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(i int) {
			fmt.Printf("진행 중 %d \n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("해결책 two 끝")
}
