package car


import "fmt"

type Bus struct {
	Charge int
}

func (b *Bus) Drive(destination string)  {
	b.Charge = 1250
	fmt.Println(destination,"까지 버스요금은  ", b.Charge, " 입니다.")
}