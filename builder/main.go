package main

import "fmt"

func main() {
	s := NewMyStruct().
		WithA(2).
		WithB(1).
		Build()

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

// Decorator .
type Decorator func(m *MyStruct) *MyStruct

// WithA .
func (f Decorator) WithA(a int) Decorator {
	return func(m *MyStruct) *MyStruct {
		f(m).a = a
		return m
	}
}

// WithB .
func (f Decorator) WithB(b int) Decorator {
	return func(m *MyStruct) *MyStruct {
		f(m).b = b
		return m
	}
}

// Build .
func (f Decorator) Build() *MyStruct {
	return f(&MyStruct{})
}

// NewMyStruct .
func NewMyStruct() Decorator {
	return func(m *MyStruct) *MyStruct {
		return m
	}
}
