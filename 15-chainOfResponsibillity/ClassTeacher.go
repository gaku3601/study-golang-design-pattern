package main

import "fmt"

type ClassTeacher struct {
	Level int
	Next  Responsibillity
}

func NewClassTeacher() *ClassTeacher {
	return &ClassTeacher{
		Level: 2,
	}
}

func (self *ClassTeacher) Judge(question *Question) {
	if self.Level >= question.Level {
		fmt.Println("担任が答えます。")
	} else if self.Next != nil {
		self.Next.Judge(question)
	} else {
		fmt.Println("わかんねｗ")
	}
}
func (self *ClassTeacher) SetNext(responsibillity Responsibillity) Responsibillity {
	self.Next = responsibillity
	return self.Next
}
