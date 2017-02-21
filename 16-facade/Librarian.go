package main

type Librarian struct{}

func NewLibrarian() *Librarian {
	return &Librarian{}
}

func (self *Librarian) searchBook(bookName string) string {
	bookList := NewBookList()
	lendingList := NewLendingList()

	for _, v := range *lendingList {
		if v.Title == bookName {
			return bookName + "は現在借りられてますね〜"
		}
	}

	for _, v := range *bookList {
		if v.Title == bookName {
			return bookName + "は現在図書館にあるけど、借りる？"
		}
	}

	return bookName + "そんな本図書館にないよ！(´・ω・｀)"
}
