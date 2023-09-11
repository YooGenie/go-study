package ageCheck

import (
	"fmt"
	"strconv"
	"time"
)

// 14세미만 체크
func CheckUnderTheAgeOfFourteen(birthDate string, seventhDigit int64) bool {
	birthMonth := birthDate[2:4] // 생월(MM)
	birthDay := birthDate[4:6]   //생일
	birthYear, _ := strconv.Atoi(birthDate[:2])

	if seventhDigit == 1 || seventhDigit == 2 || seventhDigit == 5 || seventhDigit == 6 {
		birthYear = 1900 + birthYear
	}
	if seventhDigit == 3 || seventhDigit == 4 || seventhDigit == 7 || seventhDigit == 8 {
		birthYear = 2000 + birthYear
	}

	// 생년월일 문자열을 시간 형식으로 변환
	birthDateStr := fmt.Sprintf("%d%s%s", birthYear, birthMonth, birthDay)
	birthDate8, _ := time.Parse("20060102", birthDateStr)

	if birthDate8.Format("20060102") <= time.Now().Format("20060102") {
		// 현재 시간과 비교하여 14세 미만인지 확인
		ageLimit := time.Now().AddDate(-14, 0, 0) // 현재로부터 14년 전
		birthDate8.After(ageLimit)
		return birthDate8.After(ageLimit)
	} else {
		return false
	}
}
