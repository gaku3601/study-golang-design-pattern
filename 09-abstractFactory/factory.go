package main

type Factory interface {
	getSoup() string
	getMainIngredient() string
	getVegetables() []string
}
