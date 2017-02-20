package main

import "fmt"

func main() {
	//鍋の用意
	hotPot := NewHotPot("石鍋")
	//キムチ鍋
	factory := NewKimuchiHotPotFactory()
	createHotPot(hotPot, factory)

	//味噌鍋
	factory2 := NewMisoHotPotFactory()
	createHotPot(hotPot, factory2)
}

//鍋を作成する
func createHotPot(hotPot *HotPot, factory Factory) {
	hotPot.addSoup(factory.getSoup())
	hotPot.addMainIngredient(factory.getMainIngredient())
	hotPot.addVegetables(factory.getVegetables())
	fmt.Println(hotPot)
}
