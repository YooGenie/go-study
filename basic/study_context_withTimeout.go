package basic

import (
	"context"
	"fmt"
	"time"
)

func StudyWithTimeout() {
	ctx := context.Background()
	startTime := time.Now()
	fmt.Println("시작 시간: ", startTime)
	//timeoutCtx, cancel
	timeoutCtx, _ := context.WithTimeout(ctx, time.Second*10) // 10초 후에 종료
	fmt.Println(timeoutCtx)                                   // 여기 나온 시간에 종료 된다.

	//10초 후에 타임 아웃 캔슬을 받는다.
	for {
		select {
		case <-timeoutCtx.Done():
			fmt.Println("종료 시간 : ", time.Now())
			return
		}
	}
}
