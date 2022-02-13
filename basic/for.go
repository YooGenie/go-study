package basic

import "fmt"

func StudyFor() {
	// 중첩 for 문 예시
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// continue 예시
	for i := 0; i < 5; i++ {
		if i == 0 {
			fmt.Println("i=0입니다.")
			continue
		}
		fmt.Println(i)
	}

	// break 예시
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 0 {
			fmt.Println("i=0입니다.")
			break
		}
	}
}
