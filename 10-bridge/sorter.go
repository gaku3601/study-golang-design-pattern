package main

type Sorter struct {
	sortImple SortImple
}

func NewSorter(sortImple SortImple) *Sorter {
	return &Sorter{sortImple: sortImple}
}

func (self *Sorter) sort() {
	self.sortImple.sort()
}
