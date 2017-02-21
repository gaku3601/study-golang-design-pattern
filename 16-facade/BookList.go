package main

type Book struct {
	ID    int
	Title string
}

type BookList []*Book
type LendingList []*Book

//貯蔵図書一覧を返却
func NewBookList() *BookList {
	return &BookList{
		&Book{ID: 1, Title: "昆虫図鑑"},
		&Book{ID: 2, Title: "動物図鑑"},
		&Book{ID: 3, Title: "エロ本"},
	}
}

//貸出台帳を返却
func NewLendingList() *LendingList {
	return &LendingList{
		&Book{ID: 3, Title: "エロ本"},
	}
}
