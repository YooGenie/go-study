package designPattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestBasicSingletonPatten(t *testing.T) {
	t.Run("싱글톤", func(t *testing.T) {
		s1 := BasicSingletonPatten()
		s2 := BasicSingletonPatten()
		fmt.Println("첫번째 값 : ", s1)
		fmt.Println("두번째 값 : ", s2)
		fmt.Println("동일한 값이다 : ", s1 == s2) //BasicSingletonPatten() 함수가 여러 번 호출되더라도 항상 동일한 싱글톤 인스턴스가 반환
	})
	t.Run("싱글톤-고루틴을 이용", func(t *testing.T) {
		// 두 개의 고루틴에서 같은 싱글톤 인스턴스를 사용하는 예시
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			instance1 := getInstanceByGoroutine()
			fmt.Printf("첫 번째 고루틴에서 싱글톤 인스턴스 사용: %s\n", instance1.data)
		}()

		go func() {
			defer wg.Done()
			instance2 := getInstanceByGoroutine()
			fmt.Printf("두 번째 고루틴에서 싱글톤 인스턴스 사용: %s\n", instance2.data)
		}()

		wg.Wait()
	})
}
