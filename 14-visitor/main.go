package main

func main() {
	//鈴木さん家インスタンスの作成
	suzukiHome := NewSuzukiHome()

	//鈴木さん家へ新米先生が訪問する
	rookieTeacher := NewRookieTeacher()
	suzukiHome.accept(rookieTeacher)

	//鈴木さん家へ電気工事士が訪問する
	electrial := NewElectricalContractor()
	suzukiHome.accept(electrial)
}
