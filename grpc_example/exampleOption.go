package main

import "fmt"

type Work struct {
	salary int
	place  string
}

type ExStru struct {
	age  int
	name string
	work Work
}

type Option func(*ExStru)

func WithName(inputName string) Option {
	return func(e *ExStru) {
		e.name = inputName
	}
}

func WithAge(inputSalary int, inputWorkplace string) Option {
	return func(e *ExStru) {
		e.work = Work{salary: inputSalary, place: inputWorkplace}
	}
}

func NewES(inputAge int, opts ...Option) *ExStru {
	es := &ExStru{
		age: inputAge,
	}
	for _, opt := range opts {
		opt(es)
	}
	return es
}

func main() {
	exStru := NewES(
		100,
		WithName("Meng"),
		WithAge(10000, "China"),
	)
	fmt.Println(exStru)
}
