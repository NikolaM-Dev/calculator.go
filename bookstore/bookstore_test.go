package bookstore_test

import (
	"testing"

	"github.com/NikolaM-Dev/for-the-love-of-go/bookstore"
)

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}

	want := 1

	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
}
