package designPattern

import (
	"fmt"
	"testing"
)

func TestIteratorPattern(t *testing.T) {
	t.Run("이터레이터 패턴", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		iterator := NewIntIterator(data)

		for iterator.HasNext() {
			value := iterator.Next().(int)
			fmt.Printf("현재 요소: %d\n", value)

		}
	})
}
