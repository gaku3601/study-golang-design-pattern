package main

func main() {
	//中川
	nakagawa := NewNakagawa()
	//担任
	classTeacher := NewClassTeacher()
	//ベテラン教師
	chiefTeacher := NewChiefTeacher()

	//責任の鎖を設定します。
	//中川→担任→ベテラン教師の順で答えれる問題を答えます。
	nakagawa.SetNext(classTeacher).SetNext(chiefTeacher)

	//Quetsion作成
	question1 := NewQuestion("携帯は遠足に持って行っても良いか？", 3)
	question2 := NewQuestion("バナナはおやつに入りますか？", 2)
	nakagawa.Judge(question1)
	nakagawa.Judge(question2)

	/*
		出力内容--------------
		ベテラン教師が答えます。
		担任が答えます。
	*/
}
