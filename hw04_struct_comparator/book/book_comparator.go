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
		return b1.GetYear() > b2.GetYear()
	case 2:
		return b1.GetSize() > b2.GetSize()
	case 3:
		return b1.GetRate() > b2.GetRate()
	default:
		fmt.Println("INCORRECT COMPARE!")
	}
	return false
}
