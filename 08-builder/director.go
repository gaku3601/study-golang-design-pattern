package main

type Director struct {
	Builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		Builder: builder,
	}
}

func (self *Director) constract() {
	self.Builder.addSolvent(100)
	self.Builder.addSolute(40)
	self.Builder.abandonSolution(70)
	self.Builder.addSolvent(100)
	self.Builder.addSolute(15)
	self.Builder.getResult()
}
