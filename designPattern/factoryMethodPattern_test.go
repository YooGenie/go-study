package designPattern

import (
	"fmt"
	"testing"
)

func TestFactoryMethodPattern(t *testing.T) {
	t.Run("팩토리 메소드 패턴", func(t *testing.T) {
		//클라이언트 코드에서는 구체적인 클래스가 아닌 ConcreteFactoryA를 통해 도형 객체를 생성하므로,
		//종류에 따른 객체 생성 로직을 완전히 캡슐화할 수 있습니다.
		factoryA := ConcreteFactoryA{}
		productA := factoryA.CreateProduct()
		fmt.Println("Product A:", productA.GetName())

		factoryB := ConcreteFactoryB{}
		productB := factoryB.CreateProduct()
		fmt.Println("Product B:", productB.GetName())
	})
	t.Run("팩토리 메소드 패턴- 예시", func(t *testing.T) {
		whiteShipFactory := WhiteShipFactory{}
		white := whiteShipFactory.CreateShip()
		fmt.Println("배 색깔 : ", white.Color())
		fmt.Println("배 크기 : ", white.CalculateArea())

		blockShipFactory := BlockShipFactory{}
		blockShip := blockShipFactory.CreateShip()
		fmt.Println("배 색깔 : ", blockShip.Color())
		fmt.Println("배 크기 : ", blockShip.CalculateArea())
	})
}
