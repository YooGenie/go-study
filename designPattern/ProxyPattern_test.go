package designPattern

import "testing"

func TestProxyPattern(t *testing.T) {
	t.Run("프록시 패턴", func(t *testing.T) {
		var image Image = NewProxyImage("example.jpg")

		// 이미지 로딩이 필요한 시점에서 실제 이미지가 로딩됨
		image.Display()

		// 이미지 로딩이 필요 없으므로 실제 이미지는 로딩되지 않음
		image.Display()
	})
}
