package main

type expression struct {
	operator operator
}

func newExpression(operator operator) *expression {
	return &expression{operator: operator}
}

func (e *expression) getOperandString() string {
	return e.operator.execute().getOperandString()
}
