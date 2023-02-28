package books

import "fmt"

func (b Book) PagesCount() int {
	return b.pagesCount
}

type EBook struct {
	Book
	URL string
}

//func (e *EBook) PagesCount() int {
//	return 42
//}

func DemonstrateComposition() {
	return
	eBook := EBook{
		Book: Book{
			pagesCount: 5,
		},
		URL: "test",
	}

	fmt.Println("eBook.PagesCount():", eBook.PagesCount())
	fmt.Println("eBook.Title:", eBook.Title)
}

//
//
//

type Pager struct{}

func (Pager) PagesCount() int {
	return 0
}
