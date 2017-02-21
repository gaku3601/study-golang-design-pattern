package main

import "fmt"

func main() {
	//アイスの作成
	vanilla := NewVanillaIcecream()
	greenTea := NewGreenTeaIcecream()

	fmt.Println(vanilla.getName())
	fmt.Println(vanilla.howSweet())
	fmt.Println(greenTea.getName())
	fmt.Println(greenTea.howSweet())

	//トッピング
	caVanilla := NewCashewNutsToppingIcecream(vanilla)
	caGreenTea := NewCashewNutsToppingIcecream(greenTea)

	fmt.Println(caVanilla.getName())
	fmt.Println(caVanilla.howSweet())
	fmt.Println(caGreenTea.getName())
	fmt.Println(caGreenTea.howSweet())
}
