package main

import woo "github.com/gaku3601/study-golang-design-pattern/04-templateMethod/wood"

func main() {
	wood := &woo.Wood{}
	woo.CreateWoodCutPrint(wood)
	/*出力内容------
	絵を書いて〜♪
	書いた絵を切って〜♪
	作った絵をコピーする。
	*/
}
