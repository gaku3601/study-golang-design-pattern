package main

//鍋
type HotPot struct {
	Pot            string   //鍋
	Soup           string   //スープ
	MainIngredient string   //メインの具材
	Vegetables     []string //野菜
}

//鍋を用意します。
func NewHotPot(potType string) *HotPot {
	return &HotPot{
		Pot: potType,
	}
}

//スープを加えます。
func (self *HotPot) addSoup(soupType string) {
	self.Soup = soupType
}

//メインの具材を加えます。
func (self *HotPot) addMainIngredient(ingredient string) {
	self.MainIngredient = ingredient
}

//野菜を加えます。
func (self *HotPot) addVegetables(vegetables []string) {
	self.Vegetables = vegetables
}
