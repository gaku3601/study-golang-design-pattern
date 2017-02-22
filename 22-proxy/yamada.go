package main

import "fmt"

type yamada struct{}

func newYamada() *yamada {
	return &yamada{}
}

func (y *yamada) question1() {
	fmt.Println("山田回答1")
}
func (y *yamada) question2() {
	fmt.Println("山田回答2")
}
func (y *yamada) question3() {
	fmt.Println("山田回答3")
}
