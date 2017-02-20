package main

import "fmt"

type outputText struct {
	decoStr string
}

//コンストラクタ
func NewOutputText(decoStr string) *outputText {
	return &outputText{
		decoStr: decoStr,
	}
}

func (self *outputText) use(text string) {
	fmt.Println(self.decoStr + text + self.decoStr)
}
