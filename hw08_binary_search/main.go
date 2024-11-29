package main

import "fmt"

func BinarySearch(item int, arr []int) int {
	if len(arr) < 1 {
		return -1
	}
	lowBound := 0
	highBound := len(arr) - 1

	for lowBound <= highBound {
		middle := (lowBound + highBound) / 2
		if arr[middle] == item {
			return middle
		}
		if arr[middle] > item {
			highBound = middle - 1
		} else {
			lowBound = middle + 1
		}
	}
	return -1
}

func Qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	less := []int{}
	greater := []int{}

	for i, val := range arr {
		if i == len(arr)/2 {
			continue
		}
		if val > pivot {
			greater = append(greater, val)
		} else {
			less = append(less, val)
		}
	}
	return append(append(Qsort(less), pivot), Qsort(greater)...)
}

func PrintSlice(arr []int) {
	fmt.Print("{")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println("}")
}

func main() {
	sl := []int{59, 41, 22, 8, 1, 12, 77, 99, 24, 10}
	PrintSlice(sl)
	sortedSl := Qsort(sl)
	PrintSlice(sortedSl)
	fmt.Println(BinarySearch(99, sortedSl))
}
