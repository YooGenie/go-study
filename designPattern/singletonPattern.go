package main

import (
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
	})

	return instanceByGoroutine
}

func BasicSingletonPatten() *singletonByBasic {
	if instanceByBasic == nil {
		instanceByBasic = getInstanceByBasic()
	}

	return instanceByBasic
}
