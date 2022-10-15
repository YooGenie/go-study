package basic

import (
	"context"
	"fmt"
	"time"
)

func StudyWithDeadline() {
	ctx := context.Background()
	startTime := time.Now()
	fmt.Println("시작 시간: ", startTime)
	//timeoutCtx, cancel
	deadline, _ := context.WithDeadline(ctx, time.Date(2022, 10, 15, 21, 2, 00, 00, time.Local))
	//2022년 10월 15일 21시 2분 00초 로컬 시간에 도달하면 종료

	// time.Now().Add(time.Second*10) : 현재시간에 10초 후에 종료
	fmt.Println(deadline) // 여기 나온 시간에 종료 된다.

	//지정한 날짜에 도달하면 캔슬을 받는다.
	for {
		select {
		case <-deadline.Done():
			fmt.Println("종료 신호")
			fmt.Println("종료 시간 : ", time.Now())
			return
		}
	}

}
