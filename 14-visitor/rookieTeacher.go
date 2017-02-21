package main

import "fmt"

type RookieTeacher struct{}

func NewRookieTeacher() *RookieTeacher {
	return &RookieTeacher{}
}

func (self *RookieTeacher) visit() {
	fmt.Println("こんにちわ。それでは家庭訪問を始めさせていただきたいと思います。")
}
