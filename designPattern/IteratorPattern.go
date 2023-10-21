package designPattern

// 이터레이터 인터페이스 정의 ==> 이 값만 알면 순회할 수 있다.
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// 정수 슬라이스를 순회하는 이터레이터 구현
type IntIterator struct {
	data  []int
	index int
}

func NewIntIterator(data []int) *IntIterator { // 클라이언트가 알수없다
	return &IntIterator{
		data:  data,
		index: -1,
	}
}

func (i *IntIterator) Next() interface{} {
	if i.HasNext() {
		i.index++
		return i.data[i.index]
	}
	return nil
}

func (i *IntIterator) HasNext() bool {
	return i.index+1 < len(i.data)
}
