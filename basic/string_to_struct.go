package basic

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Id            int64  `json:"id"`
	MemberId      int64  `json:"memberId"`
	ReservationNo string `json:"reservationNo"`
	NickName      string `json:"nickname"`
}

func StringToStruct() {
	str := `{
"id":100,
"memberId":1,
"nickname":"유지니",
"reservationNo":"1234567890"
}`

	structInfo := Info{}
	err := json.Unmarshal([]byte(str), &structInfo)
	if err != nil {
		fmt.Println("Unmarshal : ", err)
	}

	fmt.Println(structInfo.Id)
	fmt.Println(structInfo.NickName)
	fmt.Println(structInfo.ReservationNo)
}
