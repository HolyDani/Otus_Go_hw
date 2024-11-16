package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter size of board more than 0: ")
	fmt.Scan(&size)
	if size <= 0 {
		fmt.Println("You enter wrong value")
		return
	}

	board := GenerateBoard(size)
	fmt.Println(board)
}

func GenerateBoard(size int) string {
	if size <= 0 {
		return ""
	}

	var board string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return board
}
