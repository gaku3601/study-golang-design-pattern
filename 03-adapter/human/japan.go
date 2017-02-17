package human

type Aisatsu interface {
	Konnichiwa() string
}

type Japan struct{}

func (self *Japan) Konnichiwa() string {
	return "こんにちわ！"
}
