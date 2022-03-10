package car

import "fmt"

type Taxi struct {
	Distance int
	Onekm int
}

func (t *Taxi) Drive(destination string)  {
	charge := 3800
	fmt.Println(destination,"까지 택시요금은  ", charge+(t.Distance*t.Onekm), " 입니다.")
}

