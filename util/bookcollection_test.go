package util

import "testing"

func TestBookCount(t *testing.T) {
	type Book struct {
		ID			int `mapstructure:"id"`
		Name    	string `mapstructure:"name"`
		Author 		string `mapstructure:"author"`
		Labels      string `mapstructure:"labels"`
		Quantity 	int `mapstructure:"quantity"`
	}

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

	booklist := make([]Book, 2)
	booklist[0] = book1
	booklist[1] = book2
	testBookCollection := Collection{}
	testBookCollection.BookList = booklist
	actualCount := testBookCollection.GetCount()
	expectedCount := 2
	if actualCount != 2 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", actualCount, expectedCount)
	}
}