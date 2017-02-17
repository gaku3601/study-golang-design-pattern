package animal

import "fmt"

type Gaku struct {
	*Human
	Name string
	Age  int
}

func NewGaku() *Gaku {
	return &Gaku{Name: "gaku", Age: 25}
}

//オーバライド
func (self *Gaku) Talk() {
	fmt.Println("C#とrubyとjsとgo言語で話したい(´・ω・｀)")
}

//Gaku固有メソッド
func (self *Gaku) SelfIntroduction() {
	fmt.Printf("私は%sです。%d歳です。", self.Name, self.Age)
}
