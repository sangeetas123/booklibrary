package util

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestBookCount(t *testing.T) {

	book1 := Book {
		ID:       1,
		Name:     "book1",
		Author:   "author1",
		Labels:   "fiction",
		Quantity: 2,
	}
	book2 := Book {
		ID:       2,
		Name:     "book2",
		Author:   "author2",
		Labels:   "fiction",
		Quantity: 3,
	}

	booklist := []Book{book1, book2}
	testBookCollection := Collection{BookList: booklist}
	actualCount := testBookCollection.GetCount()
	expectedCount := 2
	assert.Equal(t, expectedCount, actualCount)
}