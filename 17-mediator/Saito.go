package main

type Saito struct{}

func NewSaito() *Saito {
	return &Saito{}
}

func (self *Saito) Consultation(colleagueInLove Colleague, secretLover Colleague) int {
	//0:片思い 1:両思い
	inLoveName := colleagueInLove.getName()
	secretName := secretLover.getName()

	//斎藤くんは誰と誰がうまくいくか知っているので、情報を返却する
	//本当はここはColleagueを使用し実装しそう...だけど、面倒なので省略
	if inLoveName == "新田" && secretName == "太郎" {
		return 1
	}
	return 0
}
