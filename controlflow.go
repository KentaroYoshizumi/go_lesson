package main

import "fmt"

//func main() {
	// if
	x := 7
	if x%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// if（短い初期化）
	if v := x * 2; v > 10 {
		fmt.Println("v > 10:", v)
	}

	// for（カウンタ）
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("sum 0..4 =", sum)

	// for（条件のみ）→ while 的に使う
	n := 3
	for n > 0 {
		n--
	}
	fmt.Println("n after loop:", n)

	// range（スライス）
	s := []int{1, 2, 3, 4, 5, 6}
	for i, val := range s {
		fmt.Println("index", i, "value", val)
	}

	// range（マップ）
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// switch（複数ケースの合流）
	switch x {
	case 1, 3, 5:
		fmt.Println("small odd")
	default:
		fmt.Println("other")
	}
}
