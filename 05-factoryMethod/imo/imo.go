package imo

import "fmt"

type Imo struct{}

func (self *Imo) Draw() {
	fmt.Println("芋に絵を書きます。")
}

func (self *Imo) Cut() {
	fmt.Println("芋を切ります。")
}
