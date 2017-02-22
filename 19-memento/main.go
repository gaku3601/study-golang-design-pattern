package main

import "fmt"

func main() {
	calc := NewCalc()
	calc.Plus(1)
	calc.Plus(1)
	calc.Plus(1)
	fmt.Println(calc.temp)

	//mementoで保管
	memento := calc.createMemento()

	//以前の計算結果を使用し、再度計算
	calc2 := NewCalc()
	calc2.setMemento(memento)
	calc2.Plus(2)
	calc2.Plus(2)
	calc2.Plus(2)
	fmt.Println(calc2.temp)
	/*
		出力結果
		3
		9
	*/
}
