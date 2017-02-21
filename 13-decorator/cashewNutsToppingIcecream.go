package main

type CashewNutsToppingIcecream struct {
	Icecream Icecream
}

func NewCashewNutsToppingIcecream(icecream Icecream) *CashewNutsToppingIcecream {
	return &CashewNutsToppingIcecream{Icecream: icecream}
}

func (self *CashewNutsToppingIcecream) getName() string {
	return "カシューナッツ" + self.Icecream.getName()
}
func (self *CashewNutsToppingIcecream) howSweet() string {
	return self.Icecream.howSweet()
}
