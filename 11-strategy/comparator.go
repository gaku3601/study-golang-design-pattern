package main

type Comparator interface {
	compare(*Human, *Human) int
}
