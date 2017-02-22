package main

func main() {
	//リモコンの用意
	remocon := newRemotoController()
	//リモコンに操作を登録する
	light := newLight()
	lightOnCommand := newLightOnCommand(light)
	lightOffCommand := newLightOffCommand(light)
	remocon.setOnCommand(lightOnCommand)
	remocon.setOffCommand(lightOffCommand)

	//現在のライトの状態を確認する
	light.displayState()

	//ライトをリモコンを使って点灯させる
	remocon.pressOnButton()

	//現在のライトの状態を確認する
	light.displayState()

	//ライトをリモコンを使って消灯させます。
	remocon.pressOffButton()

	//現在のライトの状態を確認する
	light.displayState()

	/*
		出力内容
		ライトが消えています。
		ライトが点きました。
		ライトが点いています。
		ライトが消えました。
		ライトが消えています。
	*/
}
