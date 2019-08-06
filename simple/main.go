package main

import "fmt"

func main() {
	s := NewMyStruct(1, 2)

	fmt.Println(s.PrintA())
	fmt.Println(s.PrintB())
}

// MyStruct .
type MyStruct struct {
	a int
	b int
}

// PrintA .
func (m *MyStruct) PrintA() int {
	return m.a
}

// PrintB .
func (m *MyStruct) PrintB() int {
	return m.b
}

// NewMyStruct .
func NewMyStruct(a, b int) *MyStruct {
	return &MyStruct{a, b}
}
