package main

// Book details
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}

var books = map[string]Book{
	"00": Book{Title: "Cloud Native Go", Author: "M. Leander Reimer", ISBN: "00"},
	"01": Book{Title: "Fight Club", Author: "Chuck Palahniuk", ISBN: "01"},
}

// AllBooks returns all books
func AllBooks() []Book {
	result := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		result[idx] = book
		idx++
	}

	return result
}

// GetBook returns a book given isbn
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// CreateBook creates a new book if it doesn't exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// DeleteBook removes a book given its isbn
func DeleteBook(isbn string) {
	delete(books, isbn)
}

// UpdateBook updates book and returns if update was successful
func UpdateBook(isbn string, b Book) bool {
	_, exists := books[isbn]
	if exists {
		delete(books, isbn)
		books[b.ISBN] = b
	}
	return exists
}
