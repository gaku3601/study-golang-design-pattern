package main

type VanillaIcecream struct{}

func NewVanillaIcecream() *VanillaIcecream {
	return &VanillaIcecream{}
}
func (self *VanillaIcecream) getName() string {
	return "バニラアイスクリーム"
}
func (self *VanillaIcecream) howSweet() string {
	return "バニラ味"
}
