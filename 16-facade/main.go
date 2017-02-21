package main

import "fmt"

func main() {
	librarian := NewLibrarian()
	//検索してもらう
	response := librarian.searchBook("エロ本")
	fmt.Println(response)

	response = librarian.searchBook("昆虫図鑑")
	fmt.Println(response)

	response = librarian.searchBook("動物図鑑")
	fmt.Println(response)

	response = librarian.searchBook("エロ本2")
	fmt.Println(response)

	/*
		出力結果
		エロ本は現在借りられてますね〜
		昆虫図鑑は現在図書館にあるけど、借りる？
		動物図鑑は現在図書館にあるけど、借りる？
		エロ本2そんな本図書館にないよ！(´・ω・｀)
	*/
}
