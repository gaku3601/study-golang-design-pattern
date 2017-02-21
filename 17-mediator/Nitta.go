package main

type Nitta struct {
	Tension     int
	Mediator    LoveMediator
	SecretLover Colleague
}

func NewNitta(mediator LoveMediator) *Nitta {
	return &Nitta{Mediator: mediator}
}
func (self *Nitta) getName() string {
	return "新田"
}
func (self *Nitta) setSecretLover(colleague Colleague) {
	self.SecretLover = colleague
}
func (self *Nitta) needAdvice() {
	self.Tension = self.Mediator.Consultation(self, self.SecretLover)
}
