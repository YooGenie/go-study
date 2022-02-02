package basic

import (
	"container/list"
	"fmt"
	)


type Element struct {
	Value interface{}
	Next  *Element
	Prve  *Element
}

func StudyList() {
	l := list.New()
	element4 := l.PushBack(4)
	element2 := l.PushFront(1)
	l.InsertBefore(3, element4)
	l.InsertAfter(2, element2)

	fmt.Println(l)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Next() , " ")
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	fmt.Println("From Back")

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}

