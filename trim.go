package main

import (
	"fmt"
	"strings"
)

func Trim(name string) {
	fmt.Println("----- 양쪽 공백 제거 예시 -----")
	fmt.Println("미적용 :", name)
	fmt.Println("왼쪽 trim 적용 :", strings.TrimLeft(name, " "))
	fmt.Println("오른쪽 trim 적용 :", strings.TrimRight(name, " "))
	fmt.Println("양쪽 trim 적용 :", strings.Trim(name, " "))

	fmt.Println("----- 양쪽에 다른 문자 제거 예시 -----")
	str := "유지니유"
	fmt.Println("양쪽 trim 적용 :", strings.Trim(str, "유"))
}
