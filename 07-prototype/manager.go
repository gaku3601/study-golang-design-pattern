package main

type Manager struct{}

type Product interface {
	use(string)
}

var showcase map[string]*Product = map[string]*Product{}

//prototypeの格納
func (self *Manager) register(name string, prototype Product) {
	showcase[name] = &prototype
}

//prototypeから実態を生成
func (self *Manager) create(protoName string) Product {
	return *showcase[protoName]
}
