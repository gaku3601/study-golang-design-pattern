package main

type GreenTeaIcecream struct{}

func NewGreenTeaIcecream() *GreenTeaIcecream {
	return &GreenTeaIcecream{}
}
func (self *GreenTeaIcecream) getName() string {
	return "抹茶アイスクリーム"
}
func (self *GreenTeaIcecream) howSweet() string {
	return "抹茶味"
}
