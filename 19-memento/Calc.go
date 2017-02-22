package main

type Memento struct {
	result int
}

func NewMemento(temp int) *Memento {
	return &Memento{result: temp}
}

type Calc struct {
	temp int
}

func NewCalc() *Calc {
	return &Calc{}
}

func (self *Calc) Plus(plus int) {
	self.temp += plus
}
func (self *Calc) createMemento() *Memento {
	return NewMemento(self.temp)
}
func (self *Calc) setMemento(memento *Memento) {
	self.temp = memento.result
}
