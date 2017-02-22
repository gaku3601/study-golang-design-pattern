package main

import "fmt"

type fujiwara struct {
	yamada teacher
}

func newFujiwara() *fujiwara {
	return &fujiwara{yamada: newYamada()}
}

func (f *fujiwara) question1() {
	fmt.Println("藤原回答1")
}
func (f *fujiwara) question2() {
	fmt.Println("藤原回答2")
}
func (f *fujiwara) question3() {
	//藤原先生は答えがわからないので、山田先生に聞いて回答
	f.yamada.question3()
}
