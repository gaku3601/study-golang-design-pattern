package main

type MisoHotPotFactory struct{}

func NewMisoHotPotFactory() *MisoHotPotFactory {
	return &MisoHotPotFactory{}
}

func (self *MisoHotPotFactory) getSoup() string {
	return "味噌味"
}
func (self *MisoHotPotFactory) getMainIngredient() string {
	return "鳥肉"
}
func (self *MisoHotPotFactory) getVegetables() []string {
	return []string{"白菜"}
}
