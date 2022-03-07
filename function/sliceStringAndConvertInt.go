package function

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func SliceStringAndConvertInt() {

	donationIds := "10,20,30"
	// string을 ,로 자르기
	slice := strings.Split(donationIds, ",")
	fmt.Println(slice)
	fmt.Println("타입 : ", reflect.TypeOf(slice))

	// string을 int64 변환하기
	var ArrDonationId []int64
	for _, v := range slice {
		convertToInt, _ := strconv.ParseInt(v, 10, 64)
		ArrDonationId = append(ArrDonationId, convertToInt)
	}
	fmt.Println(ArrDonationId)
	fmt.Println("타입 : ", reflect.TypeOf(ArrDonationId))

}
