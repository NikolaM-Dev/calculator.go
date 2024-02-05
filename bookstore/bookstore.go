package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     uint
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--

	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[uint]Book, ID uint) (Book, error) {
	book, ok := catalog[ID]

	if !ok {
		return Book{}, fmt.Errorf("Book with id %d, not found", ID)
	}

	return book, nil
}
