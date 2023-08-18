package main

import (
	"fmt"
)

func main() {
	// 싱글톤
	s3 := BasicSingletonPatten()
	s4 := BasicSingletonPatten()
	fmt.Println(s4, s3)
	fmt.Println(s3 == s4) //BasicSingletonPatten() 함수가 여러 번 호출되더라도 항상 동일한 싱글톤 인스턴스가 반환
}
