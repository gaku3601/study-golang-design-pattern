package main

import "fmt"

type QuickSorter struct{}

func NewQuickSorter() *QuickSorter {
	return &QuickSorter{}
}

func (self *QuickSorter) sort() {
	fmt.Println("クイックソートを実行します。")
}
