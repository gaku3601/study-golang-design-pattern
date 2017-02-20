package main

func main() {
	builder := NewSaltWaterBuilder()
	dir := NewDirector(builder)
	dir.constract()
}
