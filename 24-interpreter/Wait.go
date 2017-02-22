package main

import "strconv"

type wait struct {
	minutes int
	operand operand
}

func newWait(minutes int, operand operand) *wait {
	return &wait{minutes: minutes, operand: operand}
}
func (w *wait) execute() operand {
	return newIngredient(w.operand.getOperandString() + "を" + strconv.Itoa(w.minutes) + "分間置いたもの")
}
