package book

import "fmt"

type Compare struct {
	book1  Book
	book2  Book
	option int
}

func BuildCompare(b1, b2 *Book, opt int) *Compare {
	return &Compare{
		book1:  *b1,
		book2:  *b2,
		option: opt,
	}
}

func (comp Compare) CompareBooks() bool {
	switch comp.option {
	case 1:
		return comp.book1.GetYear() > comp.book2.GetYear()
	case 2:
		return comp.book1.GetSize() > comp.book2.GetSize()
	case 3:
		return comp.book1.GetRate() > comp.book2.GetRate()
	default:
		fmt.Println("INCORRECT COMPARE!")
	}
	return false
}
