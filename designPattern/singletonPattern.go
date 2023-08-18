package main

import (
	"fmt"
	"sync"
)

type singletonByBasic struct {
	data string
}

type singletonByGoroutine struct {
	data string
}

var (
	once                sync.Once
	instanceByBasic     *singletonByBasic
	instanceByGoroutine *singletonByGoroutine
)

func getInstanceByBasic() *singletonByBasic {
	once.Do(func() { //Do 메서드를 통해 한 번만 실행되도록 보장
		instanceByBasic = &singletonByBasic{data: "안녕"}
	})

	return instanceByBasic
}

func getInstanceByGoroutine() *singletonByGoroutine {
	once.Do(func() {
		instanceByGoroutine = &singletonByGoroutine{data: "고루틴이용"}
		fmt.Println("dddd")
	})

	return instanceByGoroutine
}

func BasicSingletonPatten() *singletonByBasic {
	if instanceByBasic == nil {
		instanceByBasic = getInstanceByBasic()
	}

	return instanceByBasic
}

func SingletonPattenByGoroutine() (*singletonByGoroutine, *singletonByGoroutine) {
	var instance1, instance2 *singletonByGoroutine
	// 두 개의 고루틴에서 같은 싱글톤 인스턴스를 사용하는 예시
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		instance1 = getInstanceByGoroutine()
		fmt.Printf("첫 번째 고루틴에서 싱글톤 인스턴스 사용: %p\n", instance1.data)
	}()

	go func() {
		defer wg.Done()
		instance2 = getInstanceByGoroutine()
		fmt.Printf("두 번째 고루틴에서 싱글톤 인스턴스 사용: %p\n", instance2.data)
	}()

	wg.Wait()

	fmt.Println(instance1 == instance2)

	return instance1, instance2
}
