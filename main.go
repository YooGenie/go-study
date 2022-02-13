package main

import (
	"fmt"
	"github.com/YooGenie/go-study/basic"
)

func main() {
	CompareIfWithSwitch(5)
	fmt.Println("")
	Trim("          유지니             ")
	fmt.Println("")
	basic.SetStruct()
	fmt.Println("")
	basic.EmbeddedField()
	fmt.Println("")
	basic.StudyIf("red", 29)
	basic.StudySwitch("sunday",28)
	basic.StudyFor()
	basic.StudyForRange()
}
