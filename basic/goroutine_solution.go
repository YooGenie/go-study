package basic

import (
	"fmt"
	"sync"
)

func GoroutineSolution() {
	fmt.Println("Goroutine 시작")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		Add()
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Goroutine 끝")
}
