package main

import (
	"fmt"

	"github.com/Otus_hw/HolyDani/hw04_struct_comparator/book"
)

func main() {
	book1 := book.NewBook(1, 1833, 224, 4.8, "Пушкин", "Евгений Онегин")
	book2 := book.NewBook(2, 1866, 672, 4.5, "Достоевский", "Преступление и наказание")
	books := book.MakeCompare(book.Rate)
	fmt.Println(books.CompareBooks(book1, book2))
}
