package models

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_Name"`
	BornYear int    `json:"born_Year"`
}

func init() {
	book1 := Book{
		ID:    1,
		Title: "Lord Of The Ring",
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkien",
			BornYear: 1892,
		},
		YearPublished: 1978,
	}

	DB = append(DB, book1)
}

func FindBookByID(id int) (Book, bool) {
	var book Book
	var found bool
	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
			break
		}
	}
	return book, found
}
