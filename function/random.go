package function

import (
	"fmt"
	"math/rand"
	"reflect"
)

func Random()  {
	randomNum := rand.Intn(1000000)
	fmt.Println("타입 체크: ",reflect.TypeOf(randomNum))
	fmt.Println("랜덤숫자: ",randomNum)
}
