package paper

type Paper struct {
	Name    string
	DrawFlg bool
	CutFlg  bool
}

func NewPaper(name string) *Paper {
	return &Paper{Name: name}
}
func (self *Paper) CreateClone() Paper {
	return *self
}
