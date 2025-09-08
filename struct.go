package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v.X = 100
	v2 := Vertex{X: 1}
	fmt.Println(v2)
}
