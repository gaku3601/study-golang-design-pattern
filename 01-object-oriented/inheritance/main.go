package main

import "github.com/gaku3601/study-golang-design-pattern/01-object-oriented/inheritance/animal"

func main() {
	//人間
	human := &animal.Human{}
	human.Work()
	human.Talk()
	//人間構造体を埋め込んだ(embed)Gaku構造体
	gaku := animal.NewGaku()
	gaku.Work()
	gaku.Talk()
	gaku.SelfIntroduction()
}
