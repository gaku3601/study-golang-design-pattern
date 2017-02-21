package main

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) *File {
	return &File{Name: name}
}

func (self *File) Remove() {
	fmt.Printf("%sを削除しました。\n", self.Name)
}
