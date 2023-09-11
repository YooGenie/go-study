package ageCheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckUnderTheAgeOfFourteen(t *testing.T) {
	t.Run("(성공) 2000년생 14세미만 검사", func(t *testing.T) {
		birthDate := "230717"
		seventhDigit := 4
		check := CheckUnderTheAgeOfFourteen(birthDate, int64(seventhDigit))

		// true : 14세 미만이다
		fmt.Println(check, "는 14세미만 입니다.")
		assert.Equal(t, true, check)
	})
	t.Run("(성공) 2000년생 14세이상 검사", func(t *testing.T) {
		birthDate := "070717"
		seventhDigit := 4
		check := CheckUnderTheAgeOfFourteen(birthDate, int64(seventhDigit))

		// true : 14세 미만이다
		fmt.Println(check, "는 14세이상 입니다.")
		assert.Equal(t, false, check)
	})
	t.Run("(성공) 1900년대 14세 이상인 경우 ", func(t *testing.T) {
		birthDate := "960717"
		seventhDigit := 1
		check := CheckUnderTheAgeOfFourteen(birthDate, int64(seventhDigit))

		fmt.Println(check, "는 14세이상 입니다.")
		assert.Equal(t, false, check)
	})
	t.Run("(실패) 현재보다 미래인 경우", func(t *testing.T) {
		birthDate := "960717"
		seventhDigit := 4
		check := CheckUnderTheAgeOfFourteen(birthDate, int64(seventhDigit))

		fmt.Println(check, "아직 태어나지 않았습니다.")
		assert.Equal(t, false, check)
	})
}
