package main

type Builder interface {
	addSolute(float64)
	addSolvent(float64)
	abandonSolution(float64)
	getResult()
}
