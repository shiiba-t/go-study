package study

import "fmt"

// Circleという構造体を定義
type Circle struct {
	Radius int
}

// Circleメソッドを定義
func (c Circle) GetArea() int {
	return 3 * c.Radius * c.Radius
}

// Squareという構造体を定義
type Square struct {
	Height int
}

// Squareメソッドを定義
func (s Square) GetArea() int {
	return s.Height * s.Height
}

// Figureインターフェースを定義
type Figure interface {
	GetArea() int
}

func DisplayArea(f Figure) {
	fmt.Printf("面積は%vです\n", f.GetArea())
}

func MainInterface() {
	circle := Circle{Radius: 1}
	DisplayArea(circle)

	square := Square{Height: 3}
	DisplayArea(square)
}
