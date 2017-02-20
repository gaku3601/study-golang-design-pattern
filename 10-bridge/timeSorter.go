package main

import "fmt"

type TimeSorter struct {
	sortImple SortImple
}

func NewTimeSorter(sortImple SortImple) *TimeSorter {
	return &TimeSorter{sortImple: sortImple}
}

func (self *TimeSorter) timeSort() {
	fmt.Println("計測を開始します。")
	self.sortImple.sort()
	fmt.Println("計測を終了します。")
}
