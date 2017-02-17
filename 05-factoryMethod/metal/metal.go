package metal

import "fmt"

type Metal struct{}

func (self *Metal) Draw() {
	fmt.Println("鉄に絵を書きます。")
}
func (self *Metal) Cut() {
	fmt.Println("鉄を切ります。")
}
