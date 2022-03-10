package car


import "fmt"

type Bus struct {

}

func (t *Bus) Drive(destination string)  {
	charge := 1250
	fmt.Println(destination,"까지 버스요금은  ", charge, " 입니다.")
}
