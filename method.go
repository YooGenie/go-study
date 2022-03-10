package main

import "fmt"

type Square struct {
	Width int64
	Vertical int64
	Total int64
}

//일반함수
func squareFunc (square *Square, vertical int64 )  {
	square.Vertical = vertical
	square.Total =  square.Vertical * square.Width
	fmt.Println("일반함수 total : ",square.Total)
}

//메서드함수
func (s *Square)squareMethod(vertical int64){
	s.Vertical = vertical
	s.Total = s.Vertical * s.Width
	fmt.Println("메서드 total : ",s.Total)
}

// 리시버가 있으면 메서드입니다. 리시버는 메서드가 속한 타입을 알려줍니다.
// methodType이라는 메서드는 square라는 구조체에 속합니다.

// 메서드를 이용하는 이유는 코드 응집도를 높이기 위해서 입니다. 메서드는 데이터와 기능을 묶습니다.


func Method()  {
	sPointer := &Square{Width: 100} //var s1 = new(Square)
	squareFunc(sPointer, 30)
	fmt.Println("가로 :",sPointer.Width,"세로 : ",sPointer.Vertical, "total 값 : ",sPointer.Total)
	sPointer.squareMethod(30)
	fmt.Println("가로 :",sPointer.Width,"세로 : ",sPointer.Vertical, "total 값 : ",sPointer.Total)
}

//메서드와 함수는 표현하는 방법이 다르다.  methodType메서드는 리시버 타입이 square인 메서드입니다.
//메서드는 구조체와 비슷하게 s. 점으로 호출합니다.



