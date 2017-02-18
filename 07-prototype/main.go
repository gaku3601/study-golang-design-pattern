package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/07-prototype/paper"
)

func main() {
	prototype := paper.NewPaper("雪の結晶")
	draw(prototype)
	cut(prototype)
	fmt.Println(prototype)
	//ここまででプロトタイプ完成

	//複製処理
	paper1 := prototype.CreateClone()
	paper1.Name = "雪の結晶2"
	paper2 := prototype.CreateClone()
	paper2.Name = "雪の結晶3"

	fmt.Println(paper1)
	fmt.Println(paper2)
}

func draw(paper *paper.Paper) {
	//時間のかかる処理
	paper.DrawFlg = true
}

func cut(paper *paper.Paper) {
	//時間のかかる処理
	paper.CutFlg = true
}
