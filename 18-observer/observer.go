package main

import "fmt"

type Observer interface {
	update(string)
}

type Teacher1 struct{}
type Teacher2 struct{}

func NewTeacher1() *Teacher1 {
	return &Teacher1{}
}
func NewTeacher2() *Teacher2 {
	return &Teacher2{}
}

func (t *Teacher1) update(message string) {
	fmt.Println(message)
}
func (t *Teacher2) update(message string) {
	fmt.Println(message)
}
