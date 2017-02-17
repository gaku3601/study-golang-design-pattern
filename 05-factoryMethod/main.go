package main

import (
	im "github.com/gaku3601/study-golang-design-pattern/05-factoryMethod/imo"
	met "github.com/gaku3601/study-golang-design-pattern/05-factoryMethod/metal"
	woo "github.com/gaku3601/study-golang-design-pattern/05-factoryMethod/wood"
)

func main() {
	wood := &woo.Wood{}
	processing(wood)
	metal := &met.Metal{}
	processing(metal)
	imo := &im.Imo{}
	processing(imo)
}
