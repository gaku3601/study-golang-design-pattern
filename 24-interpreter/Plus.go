package main

type plus struct {
	operand1 operand
	operand2 operand
}

func newPlus(operand1 operand, operand2 operand) *plus {
	return &plus{operand1: operand1, operand2: operand2}
}

func (p *plus) execute() operand {
	return newIngredient(p.operand1.getOperandString() + "と" + p.operand2.getOperandString() + "を足したもの")
}
