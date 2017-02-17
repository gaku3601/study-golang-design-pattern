package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/03-adapter/human"
)

func main() {
	japan := &human.Japan{}
	output(japan)

	english := &human.English{}
	output(english)
}

//ここはテスト済みだから変えたくない
func output(aisatsu human.Aisatsu) {
	fmt.Println(aisatsu.Konnichiwa())
}
