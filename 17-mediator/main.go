package main

import "fmt"

func main() {
	saito := NewSaito()

	nitta := NewNitta(saito)
	taro := NewTaro(saito)

	//片思いをセット
	nitta.setSecretLover(taro)
	//斎藤くんに相談して、両思いか聞いてみる
	nitta.needAdvice()
	fmt.Println(nitta.Tension)
}
