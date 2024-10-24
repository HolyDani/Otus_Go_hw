package main

import "fmt"

func main() {
	var size int
	var board string
	fmt.Println("Enter size of board more than 0: ")
	fmt.Scan(&size)
	if size <= 0 {
		fmt.Println("You enter wrong value")
		return
	}

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

	fmt.Println(board)
}
