package main

type Human struct {
	Name   string
	Height int
	Weight int
	Age    int
}

func NewHuman(name string, height int, weight int, age int) *Human {
	return &Human{
		Name:   name,
		Height: height,
		Weight: weight,
		Age:    age,
	}
}
