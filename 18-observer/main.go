package main

func main() {
	//通知先の先生を登録
	t1 := NewTeacher1()
	t2 := NewTeacher2()

	//生徒の作成
	s := NewStudent()

	//通知先の設定
	s.addObserver(t1)
	s.addObserver(t2)

	//通知する
	s.notifyObservers()
}
