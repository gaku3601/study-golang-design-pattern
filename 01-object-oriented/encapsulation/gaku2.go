package main

type Gaku2 struct {
	Name string
	Age  int
}

func NewGaku2() *Gaku2 {
	return &Gaku2{Name: "gaku", Age: 25}
}

func (self *Gaku2) AfterTenYears() {
	self.Age += 10
}
func (self *Gaku2) beforeTenYears() {
	self.Age -= 10
}
