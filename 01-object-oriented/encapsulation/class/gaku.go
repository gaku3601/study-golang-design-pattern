package class

type Gaku struct {
	Name string
	Age  int
}

func NewGaku() *Gaku {
	return &Gaku{Name: "gaku", Age: 25}
}

func (self *Gaku) AfterTenYears() {
	self.Age += 10
}
func (self *Gaku) beforeTenYears() {
	self.Age -= 10
}
