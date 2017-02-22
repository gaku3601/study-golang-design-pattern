package main

type BadMoodState struct{}

func NewBadMoodState() *BadMoodState {
	return &BadMoodState{}
}
func (b *BadMoodState) morningGreet() string {
	return "話しかけてくんな！"
}
func (b *BadMoodState) getProtectionForCold() string {
	return "寒いから、話しかけてくんな!"
}
