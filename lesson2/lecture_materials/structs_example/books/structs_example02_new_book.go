package books

import (
	"errors"
)

func New(author, title string, pages []string) (Book, error) {
	if title == "" {
		return Book{}, errors.New("empty book title is not allowed")
	}

	return Book{
		Author:     author,
		Title:      title,
		pagesCount: len(pages),
	}, nil
}
