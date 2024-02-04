package bookstore_test

import (
	"testing"

	"github.com/NikolaM-Dev/for-the-love-of-go/bookstore"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}
