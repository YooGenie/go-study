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
	byteContent, _ := json.Marshal(str)

	err := json.Unmarshal(byteContent, &structInfo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(structInfo.Id)
}
