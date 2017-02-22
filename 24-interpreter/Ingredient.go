package main

type ingredient struct {
	operandString string
}

func newIngredient(operandString string) *ingredient {
	return &ingredient{operandString: operandString}
}

func (i *ingredient) getOperandString() string {
	return i.operandString
}
