package main

type IMaterial interface {
	Draw()
	Cut()
}

func processing(material IMaterial) {
	material.Draw()
	material.Cut()
}
