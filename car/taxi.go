package car

import "fmt"

type Taxi struct {
	baseCost int
	Distance int
	Onekm int
}


func (t *Taxi) Drive(destination string)  {
	t.baseCost= 3800
	fmt.Println(destination,"까지 택시요금은  ", t.baseCost+(t.Distance*t.Onekm), " 입니다.")
}
