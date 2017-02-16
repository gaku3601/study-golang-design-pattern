package class

type Gaku struct {
	Name string
	Age  int
}

//コンストラクタ
func NewGaku(name string, age int) *Gaku {
	return &Gaku{Name: name, Age: age}
}
