package main

import "fmt"

type Nakagawa struct {
	Level int
	Next  Responsibillity
}

func NewNakagawa() *Nakagawa {
	return &Nakagawa{Level: 1}
}

func (self *Nakagawa) Judge(question *Question) {
	if self.Level >= question.Level {
		fmt.Println("中川が答えます。")
	} else if self.Next != nil {
		self.Next.Judge(question)
	} else {
		fmt.Println("わかんねｗ")
	}
}
func (self *Nakagawa) SetNext(responsibillity Responsibillity) Responsibillity {
	self.Next = responsibillity
	return self.Next
}
