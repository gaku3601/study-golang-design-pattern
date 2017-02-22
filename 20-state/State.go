package main

type State interface {
	morningGreet() string
	getProtectionForCold() string
}
