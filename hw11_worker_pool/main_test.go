package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	var counter int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	expected := 1000000

	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	assert.Equal(t, expected, counter)
}
