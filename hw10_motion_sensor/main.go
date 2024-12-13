package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func RandIntPush(buf chan<- int) {
	t := time.After(60 * time.Second)
wait:
	for {
		select {
		case <-t:
			close(buf)
			break wait
		default:
			n, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				fmt.Println("Error generating random number:", err)
				continue
			}
			buf <- int(n.Int64())
		}
	}
}

func Average(ch <-chan int, result chan int) {
	var sum, count int
	for item := range ch {
		sum += item
		count++
		if count == 10 {
			result <- sum / 10
			sum = 0
			count = 0
		}
	}
	if count > 0 {
		result <- sum / count
	}
	close(result)
}

func main() {
	processed := make(chan int)
	buff := make(chan int)

	go Average(buff, processed)
	go RandIntPush(buff)

	for num := range processed {
		fmt.Println("Average:", num)
	}
}
