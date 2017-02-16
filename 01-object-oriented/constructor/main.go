package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/01-object-oriented/constructor/class"
)

func main() {
	//インスタンス生成
	gaku := class.NewGaku("gaku", 25)
	fmt.Println(gaku.Name, gaku.Age)
	// 出力:gaku 25
}
