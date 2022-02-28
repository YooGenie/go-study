package main

import (
	"fmt"
	"github.com/YooGenie/go-study/basic"
	"github.com/YooGenie/go-study/compare"
	"github.com/YooGenie/go-study/function"
)

func main() {
	compare.CompareIfWithSwitch(5)
	fmt.Println("")
	function.Trim("          유지니             ")
	fmt.Println("")
	basic.SetStruct()
	fmt.Println("")
	basic.EmbeddedField()
	fmt.Println("")
	basic.StudyIf("red", 29)
	basic.StudySwitch("sunday",28)
	basic.StudyFor()
	basic.StudyForRange()
	basic.StudyArray()
	function.CheckType()
	basic.SliceStudy()
	function.FillZero(12365)
}
