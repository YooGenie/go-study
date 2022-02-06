package main

import "fmt"

type Donor struct {
	DonationId      int64
	NationalityType string
	DonorType       string
	Name            string
	Mobile          string
	RegistrationNo  string
	Email           string
	Check           bool
}

type DonorReceipt struct {
	Id     int64
	Status string
	Type   string
	Note   string
	Name   string
	Donor  Donor
}

func SetStruct() {
	var donorStruct Donor
	fmt.Println("초기값 설정 전 :", donorStruct)

	donorStruct.DonationId = 123
	donorStruct.Email = "1@test.com"
	donorStruct.NationalityType = "NATIVE"
	donorStruct.DonorType = "INDIVIDUAL"
	donorStruct.Name = "홍길동"
	donorStruct.Mobile = "0000000000"
	donorStruct.Check = true

	fmt.Println("초기값 설정 후 :", donorStruct)
	fmt.Println("이름 필드 값 출력 :", donorStruct.Name)

	//모든 필드 초기화
	var donorStructV1 = Donor{1, "NATIVE", "INDIVIDUAL", "홍길동", "0000000000", "9901012000000", "1@test.com", true}
	fmt.Println("V1 초기값 설정 후 :", donorStructV1)

	var donorStructV2 = Donor{
		1,
		"NATIVE",
		"INDIVIDUAL",
		"홍길동",
		"0000000000",
		"9901012000000",
		"1@test.com",
		true,
	}
	fmt.Println("V2 초기값 설정 후 :", donorStructV2)

	//일부 필드만 초기하는 방법
	var donorStructV3 = Donor{Name: "홍길동", Email: "1@test.com"}
	fmt.Println("V3 초기값 설정 후 :", donorStructV3)

	var donorStructV4 = Donor{
		DonationId: 100,
		Name:       "홍길동",
		Email:      "1@test.com",
	}
	fmt.Println("V4 초기값 설정 후 :", donorStructV4)
}

func EmbeddedField() {
	var donor = Donor{
		1,
		"NATIVE",
		"INDIVIDUAL",
		"홍길동",
		"0000000000",
		"9901012000000",
		"1@test.com",
		true,
	}

	var donorReceiptV1 = DonorReceipt{
		Id:     1,
		Status: "Applied",
		Donor:  donor,
	}
	fmt.Println("donor 초기값 설정 후 :", donor)
	fmt.Println("donorReceipt V1 초기값 설정 후 :", donorReceiptV1)

	var donorReceiptV2 = DonorReceipt{
		Id:     1,
		Status: "Applied",
		Donor: Donor{
			DonorType: "INDIVIDUAL",
			Name:      "김철수",
		},
	}
	fmt.Println("donorReceipt V2 초기값 설정 후 :", donorReceiptV2)

	var donorReceipt = DonorReceipt{
		Id:     1,
		Status: "Applied",
		Name:   "김영희",
		Donor: Donor{
			DonorType: "INDIVIDUAL",
			Name:      "김철수",
		},
	}
	fmt.Println("donorReceipt 필드안에 있는 donor의 name 값 :", donorReceipt.Donor.Name)
	fmt.Println("donorReceipt 안에 name 값 :", donorReceipt.Name)
}
