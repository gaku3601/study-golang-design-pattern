package main

type Yumi struct {
	state State
}

func NewYumi(state State) *Yumi {
	return &Yumi{state: state}
}
func (y *Yumi) morningGreet() string {
	return y.state.morningGreet()
}
func (y *Yumi) getProtectionForCold() string {
	return y.state.getProtectionForCold()
}
