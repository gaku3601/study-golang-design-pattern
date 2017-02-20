package main

import "fmt"

type doubleText struct {
	decoStr string
}

func NewDoubleText(decoStr string) *doubleText {
	return &doubleText{
		decoStr: decoStr,
	}
}

//2回表示する
func (self *doubleText) use(text string) {
	fmt.Println(self.decoStr + text + self.decoStr)
	fmt.Println(self.decoStr + text + self.decoStr)
}
