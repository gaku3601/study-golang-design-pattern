package singleton

type single struct {
	Name string
}

var sharedInstance *single = newSingle()

func newSingle() *single {
	return &single{Name: "gaku"}
}

func GetInstance() *single {
	return sharedInstance
}
func (self *single) ReName(name string) {
	self.Name = name
}
