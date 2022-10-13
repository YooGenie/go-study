package basic

import (
	"context"
	"fmt"
)

func StudyWithCancel() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	cancel()

	// 1
	fmt.Println("종료 전")
	select {
	case <-ctx.Done():
		fmt.Println("종료 시그널")
		return
	}

	//2
	fmt.Println(ctx.Err()) //nil
	cancel()               // 호출하면 컨텍스트가 종료가 된다.
	fmt.Println(ctx.Err()) //context canceled

}
