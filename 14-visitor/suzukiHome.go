package main

type SuzukiHome struct{}

func NewSuzukiHome() *SuzukiHome {
	return &SuzukiHome{}
}

func (self *SuzukiHome) accept(visitor Visitor) {
	visitor.visit()
}
