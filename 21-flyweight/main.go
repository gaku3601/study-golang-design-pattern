package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/21-flyweight/HumanLetterFactory"
)

func main() {
	factory := HumanLetterFactory.GetInstance()
	a1 := factory.GetHumanLetter("あ")
	a2 := factory.GetHumanLetter("あ")
	a3 := factory.GetHumanLetter("あ")
	fmt.Println(a1.Letter)
	fmt.Println(a2.Letter)
	fmt.Println(a3.Letter)
	i1 := factory.GetHumanLetter("い")
	fmt.Println(i1.Letter)
	fmt.Println(a1.Letter)

	factory2 := HumanLetterFactory.GetInstance()
	u1 := factory2.GetHumanLetter("う")
	fmt.Println(u1.Letter)

}
