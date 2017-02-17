package animal

import "fmt"

type Human struct{}

//人間は歩きます
func (self *Human) Work() {
	fmt.Println("歩く")
}

//人間は喋ります
func (self *Human) Talk() {
	fmt.Println("しゃべる")
}
