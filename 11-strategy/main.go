package main

import "fmt"

func main() {
	//人間の作成
	human1 := NewHuman("勇者gaku", 170, 65, 28)
	human2 := NewHuman("ニートgaku", 170, 72, 25)

	//比較インスタンスの作成
	ageComparator := &AgeComparator{}

	//比較結果算出
	com := NewMyClass(ageComparator)
	result := com.compare(human1, human2)
	fmt.Println(result)
}
