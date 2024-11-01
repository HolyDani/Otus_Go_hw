package book

import "fmt"

type Compare struct {
	option int
}

func MakeCompare(opt int) Compare {
	return Compare{
		option: opt,
	}
}

func (comp Compare) CompareBooks(b1, b2 *Book) bool {
	switch comp.option {
	case 1:
		return *b1.Year() > *b2.Year()
	case 2:
		return *b1.Size() > *b2.Size()
	case 3:
		return *b1.Rate() > *b2.Rate()
	default:
		fmt.Println("INCORRECT COMPARE!")
	}
	return false
}
