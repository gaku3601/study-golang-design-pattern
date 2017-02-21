package main

type Taro struct {
	Tension     int
	Mediator    LoveMediator
	SecretLover Colleague
}

func NewTaro(mediator LoveMediator) *Taro {
	return &Taro{Mediator: mediator}
}
func (self *Taro) getName() string {
	return "太郎"
}
func (self *Taro) setSecretLover(colleague Colleague) {
	self.SecretLover = colleague
}
func (self *Taro) needAdvice() {
	self.Tension = self.Mediator.Consultation(self, self.SecretLover)
}
