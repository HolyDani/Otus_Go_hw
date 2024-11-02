package book

import "fmt"

type State int

const (
	Year State = iota + 1
	Size
	Rate
)

type Compare struct {
	option State
}

func MakeCompare(opt State) Compare {
	return Compare{
		option: opt,
	}
}

func (comp Compare) CompareBooks(b1, b2 *Book) bool {
	switch comp.option {
	case Year:
		return *b1.Year() > *b2.Year()
	case Size:
		return *b1.Size() > *b2.Size()
	case Rate:
		return *b1.Rate() > *b2.Rate()
	default:
		fmt.Println("INCORRECT COMPARE!")
	}
	return false
}
