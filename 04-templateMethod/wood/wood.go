package wood

import "fmt"

type Wood struct{}

func (self *Wood) draw() {
	fmt.Println("絵を書いて〜♪")
}
func (self *Wood) cut() {
	fmt.Println("書いた絵を切って〜♪")
}
func (self *Wood) coppy() {
	fmt.Println("作った絵をコピーする。")
}
