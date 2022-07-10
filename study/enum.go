package study

import "fmt"

type Color int

// Go言語におけるenum(列挙型)の書き方
// iota(イオタ)はconst内で使用される、型なしの連続する整数定数
// 0から始まり、各定数定義の後に1ずつインクリメントされる
// 2つ目以降の定数はiota省略可能
const (
	Red    Color = iota // Red    == 0
	Blue                // Ble    == 1
	Yellow              // Yellow == 2
)

// Stringインターフェースを実装
func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	default:
		return ""
	}
}

func StudyEnum() {
	fmt.Println(Red)
}
