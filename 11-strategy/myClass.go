package main

type MyClass struct {
	comparator Comparator
}

func NewMyClass(comparator Comparator) *MyClass {
	return &MyClass{comparator: comparator}
}

func (self *MyClass) compare(human1 *Human, human2 *Human) int {
	return self.comparator.compare(human1, human2)
}
