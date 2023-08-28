package designPattern

import "fmt"

// Subject 인터페이스 정의 => Image 인터페이스는 Display() 메서드를 가지고 있다.
type Image interface {
	Display()
}

// RealImage는 실제 이미지를 나타내는 구조체
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Loading image:", filename)
	return &RealImage{filename: filename}
}

func (r *RealImage) Display() {
	fmt.Println("Displaying image:", r.filename)
}

// ProxyImage는 이미지 로딩을 지연시키는 프록시 구조체
type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

// ProxyImage의 Display() 메서드에서는 실제 이미지가 로딩되기 전에는 realImage가 초기화되지 않습니다.
func (p *ProxyImage) Display() {
	if p.realImage == nil { // 이미지가 실제로 필요한 순간에 realImage를 생성하고 로딩하여 실제 이미지를 보여줍니다.
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display() // 실제 이미지를 표시
}
