package function

import (
	"fmt"
	"strings"
)

func Join() {
	var nums = []string{"3", "5", "6", "1", "0", "66"}
	str := strings.Join(nums,"")
	fmt.Println(str)

	var arr = []string{"오렌지", "사과", "포도", "귤", "바나나", "망고"}
	arrStr := strings.Join(arr,",")
	fmt.Println(arrStr)

}