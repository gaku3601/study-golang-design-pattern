package main

type Responsibillity interface {
	Judge(*Question)
	SetNext(Responsibillity) Responsibillity
}
