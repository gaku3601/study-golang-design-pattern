package wood

type WoodCutPrint interface {
	draw()
	cut()
	coppy()
}

func CreateWoodCutPrint(woodCutPrint WoodCutPrint) {
	//絵を書きます
	woodCutPrint.draw()
	//書いた絵を切ります
	woodCutPrint.cut()
	//絵をcoppyします。
	woodCutPrint.coppy()
}
