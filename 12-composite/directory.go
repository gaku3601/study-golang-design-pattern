package main

import "fmt"

type Directory struct {
	List []FileAndDir
	Name string
}

func NewDirectory(name string) *Directory {
	return &Directory{Name: name}
}

func (self *Directory) Add(fileAndDir FileAndDir) {
	self.List = append(self.List, fileAndDir)
}

func (self *Directory) Remove() {
	for _, v := range self.List {
		v.Remove()
	}
	fmt.Printf("%sを削除しました。\n", self.Name)
}
