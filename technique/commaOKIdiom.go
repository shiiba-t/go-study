package technique

import "fmt"

// パターン１(map探索)
func Comma1() {
	// mapの作成
	f := map[string]int{
		"Apple":  1,
		"Banana": 2,
	}

	// mapの中に指定の情報が存在しているかを確認
	// 存在すればokにはtrueが入る
	if s, ok := f["Apple"]; ok {
		// s=1, ok=true
		fmt.Printf("s=%v ok=%v \n", s, ok)
	}

	// Orangeはマップの中にないので、if文の中は実行されない
	if s, ok := f["Orange"]; ok {
		// not execute
	} else {
		// s=0, ok=false
		fmt.Printf("s=%v ok=%v \n", s, ok)
	}
}

// パターン2(型チェック)
func Comma2() {
	var test interface{}
	test = 123

	if t, ok := test.(int); ok {
		// s=123, ok=true
		fmt.Printf("t=%v ok=%v \n", t, ok)
	}

	if t, ok := test.(string); ok {
		// not execute
	} else {
		// t="", ok=false
		fmt.Printf("t=%v ok=%v \n", t, ok)
	}

}
