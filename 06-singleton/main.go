package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/06-singleton/singleton"
)

func main() {
	single := singleton.GetInstance()
	fmt.Println(single.Name)
	single.ReName("gaku2")
	fmt.Println(single.Name)
	single = singleton.GetInstance()
	fmt.Println(single.Name)
}
