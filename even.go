package main // ① package 宣言は必ずファイルの一番上

import "fmt" // ② import は package 宣言の直後

//func main() { // ③ 関数の中に処理を書く
	for i := 1; i <= 30; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
