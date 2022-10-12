package basic

import (
	"context"
	"fmt"
)

func StudyContext() {
	//초기화를 해준다
	ctx := context.Background()
	//context.TODO()
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println(ctx.Value("key")) //value
	fmt.Println(ctx.Value("no"))  // <nil>

}
