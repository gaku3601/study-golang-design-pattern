package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/01-object-oriented/encapsulation/class"
)

func main() {
	gaku := class.NewGaku()
	gaku.AfterTenYears()
	fmt.Println(gaku.Age)

	//privateになっているのでエラーとなる
	//gaku.beforeTenYears()
	//fmt.Println(gaku.Age)

	//ちなみに同パッケージ内であればprivateは適用されないので注意
	gaku2 := NewGaku2()
	gaku2.AfterTenYears()
	fmt.Println(gaku2.Age)

	gaku2.beforeTenYears()
	fmt.Println(gaku2.Age)
}
