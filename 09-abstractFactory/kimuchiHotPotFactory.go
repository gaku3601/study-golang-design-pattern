package main

type KimuchiHotPotFactory struct{}

func NewKimuchiHotPotFactory() *KimuchiHotPotFactory {
	return &KimuchiHotPotFactory{}
}

func (self *KimuchiHotPotFactory) getSoup() string {
	return "キムチ味"
}
func (self *KimuchiHotPotFactory) getMainIngredient() string {
	return "豚肉"
}
func (self *KimuchiHotPotFactory) getVegetables() []string {
	return []string{"ニラ", "白菜", "ニラ"}
}
