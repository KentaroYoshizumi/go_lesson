package main

import "fmt"

//func main() {
	// 宣言パターン
	var a int = 10 // 明示的に型を付ける
	)b := 20        // 短縮宣言（関数内のみ）
	var s string = "hello"
	var flag bool       // bool のゼロ値は false
	const Pi = 3.141592 // 定数

	fmt.Println("a b s flag Pi:", a, b, s, flag, Pi)

	// 型変換（異種の数値を混ぜるときは必須）
	sum := float64(a) + Pi
	fmt.Printf("sum: %.3f\n", sum)

	// 複数代入と型推論
	var x, y = 1, 2.5
	fmt.Printf("x=%v %T, y=%v %T\n", x, x, y, y)

	// ゼロ値の確認
	var zeroInt int
	var zeroStr string
	fmt.Printf("zeroInt=%d, zeroStr=%q\n", zeroInt, zeroStr)
}
