package main

import "fmt"

func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// ss := i.(string)
	// fmt.Println(ss + "!")
	// i.(type)
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	case bool:
		fmt.Println(v)
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	// var i interface{} = 10
	do(10)
	do("Mike")
	do(true)
	do(3.14)
	// fmt.Println(do())
}
