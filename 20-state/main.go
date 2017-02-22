package main

import "fmt"

func main() {
	//機嫌の悪い由美の作成
	bad := NewBadMoodState()
	badYumi := NewYumi(bad)
	fmt.Println(badYumi.morningGreet())
	fmt.Println(badYumi.getProtectionForCold())

	//機嫌の良い由美の作成
	good := NewGoodState()
	goodYumi := NewYumi(good)
	fmt.Println(goodYumi.morningGreet())
	fmt.Println(goodYumi.getProtectionForCold())

	/*
		出力結果
		話しかけてくんな！
		寒いから、話しかけてくんな!
		今日は良い天気ですね。
		寒いね〜
	*/
}
