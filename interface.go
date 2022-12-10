package main

import "github.com/YooGenie/go-study/car"

//인터페이스: 구현을 포함하지 않은 메서드 집합입니다.

// 인터페이스 선언
type InterfaceName interface {
	//{} 안에 메서드 집합
}

type Car interface {
	Drive(donor string)
}

func Person(destination string, c Car) {
	c.Drive(destination) //Car 인터페이스로 Drive 메서드 호출

}

func Interface() {

	taxi := &car.Taxi{Distance: 10, Onekm: 100}
	Person("강남", taxi)

	bus := &car.Bus{}
	Person("강남", bus)

}

// 덕타이핑
// 어떤 타입이 인터페이스에 포함하고 있는지 여부를 결정할 때 덕 타이핑을 사용한다.
//인터페이스에 정의한 메서드를 포함 여부만 결정하는 방식
