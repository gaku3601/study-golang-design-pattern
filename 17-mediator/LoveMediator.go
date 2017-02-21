package main

type LoveMediator interface {
	Consultation(Colleague, Colleague) int
}
