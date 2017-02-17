package wood

import "fmt"

type Wood struct{}

func (self *Wood) Draw() {
	fmt.Println("木に絵を書きます。")
}

func (self *Wood) Cut() {
	fmt.Println("木を切ります。")
}
