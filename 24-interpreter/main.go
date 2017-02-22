package main

import "fmt"

func main() {
	//カップ麺の材料を作成する
	kona := newIngredient("粉")
	yu := newIngredient("湯")
	men := newIngredient("麺")

	//処理を行う
	konaAndMen := newPlus(kona, men).execute()
	konaAndMenAndYu := newPlus(konaAndMen, yu).execute()
	konaAndMenAndYuWait3 := newWait(3, konaAndMenAndYu)

	expression := newExpression(konaAndMenAndYuWait3)
	fmt.Println(expression.getOperandString())
	/*
		出力結果
		粉と麺を足したものと湯を足したものを3分間置いたもの
	*/
}
