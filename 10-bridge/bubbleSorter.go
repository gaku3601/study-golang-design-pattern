package main

import "fmt"

type BubbleSorter struct{}

func NewBubbleSorter() *BubbleSorter {
	return &BubbleSorter{}
}

func (self *BubbleSorter) sort() {
	fmt.Println("バブルソートを実行します。")
}
