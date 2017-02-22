package main

type GoodState struct{}

func NewGoodState() *GoodState {
	return &GoodState{}
}
func (b *GoodState) morningGreet() string {
	return "今日は良い天気ですね。"
}
func (b *GoodState) getProtectionForCold() string {
	return "寒いね〜"
}
