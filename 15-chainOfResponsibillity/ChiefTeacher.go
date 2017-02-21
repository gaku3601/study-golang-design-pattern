package main

import "fmt"

type ChiefTeacher struct {
	Level int
	Next  Responsibillity
}

func NewChiefTeacher() *ChiefTeacher {
	return &ChiefTeacher{Level: 3}
}

func (self *ChiefTeacher) Judge(question *Question) {
	if self.Level >= question.Level {
		fmt.Println("ベテラン教師が答えます。")
	} else if self.Next != nil {
		self.Next.Judge(question)
	} else {
		fmt.Println("わかんねｗ")
	}
}

func (self *ChiefTeacher) SetNext(responsibillity Responsibillity) Responsibillity {
	self.Next = responsibillity
	return self.Next
}
