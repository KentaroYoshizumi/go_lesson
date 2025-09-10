package main

import "fmt"

// 共通の構造体
type Person struct {
	Name string
}

// Personを埋め込む
type Student struct {
	Person
	School string
}

func main() {
	s := Student{
		Person: Person{Name: "Alice"},
		School: "Go University",
	}

	// 埋め込みのおかげで s.Person.Name を省略できる
	fmt.Println(s.Name, "studies at", s.School)
}
