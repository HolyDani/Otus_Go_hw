package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	input := make(chan int)
	output := make(chan int)

	go Average(input, output)

	for i := 1; i <= 10; i++ {
		input <- i
	}
	close(input)

	result, ok := <-output
	expectedResult := 55 / 10
	expectedStatus := true
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedStatus, ok)
}

func TestRandIntPush(t *testing.T) {
	buf := make(chan int)

	go RandIntPush(buf)

	time.Sleep(3 * time.Second)

	_, ok := <-buf
	expectedStatus := true
	assert.Equal(t, expectedStatus, ok)
}
