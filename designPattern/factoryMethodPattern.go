package designPattern

// Product 인터페이스
type Product interface { // 공통으로 작동하는 걸 적는다.
	GetName() string
}

// ConcreteProductA 구현
type ConcreteProductA struct{}

func (cpa ConcreteProductA) GetName() string {
	return "Concrete Product A"
}

// ConcreteProductB 구현
type ConcreteProductB struct{}

func (cpb ConcreteProductB) GetName() string {
	return "Concrete Product B"
}

// 팩토리 인터페이스
type Factory interface { //객체를 생성하는 메소드를 정의
	CreateProduct() Product
}

// 팩토리마다 CreateProduct를

// ConcreteFactoryA 구현
type ConcreteFactoryA struct{}

func (cfa ConcreteFactoryA) CreateProduct() Product {
	return ConcreteProductA{}
}

// ConcreteFactoryB 구현
type ConcreteFactoryB struct{}

func (cfb ConcreteFactoryB) CreateProduct() Product {
	return ConcreteProductB{}
}

// 새로운 예시

// Ship 인터페이스
type Ship interface { // 공통 동작을 정의
	Color() string // 공통기능
	CalculateArea() float64
}

type WhiteShip struct {
	Size float64
}

func (c WhiteShip) Color() string {
	return "color : white"
}

func (c WhiteShip) CalculateArea() float64 {
	return c.Size * c.Size
}

type BlockShip struct {
	Length float64
}

func (r BlockShip) Color() string {
	return "color : block"
}

func (r BlockShip) CalculateArea() float64 {
	return r.Length
}

// 팩토리 인터페이스
type ShipFactory interface {
	CreateShip() Ship
}

// WhiteShipFactory 구현
type WhiteShipFactory struct{}

func (cf WhiteShipFactory) CreateShip() Ship {
	return WhiteShip{Size: 5}
}

// BlockShipFactory 구현
type BlockShipFactory struct{}

func (rf BlockShipFactory) CreateShip() Ship {
	return BlockShip{Length: 500}
}
