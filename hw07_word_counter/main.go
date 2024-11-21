package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrEmptyString = errors.New("empty string")

func main() {
	text := "Hello my friend. You are my best friend!"
	words, err := countWords(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range words {
		fmt.Printf("Word: %s = %d\n", key, value)
	}
}

func countWords(str string) (map[string]uint, error) {
	re := regexp.MustCompile(`[[:punct:]]`)
	cleanedText := re.ReplaceAllString(str, "")
	if len(cleanedText) > 0 && strings.TrimSpace(cleanedText) != "" {
		wordCount := make(map[string]uint)
		cleanedText = strings.ToLower(cleanedText)
		words := strings.Fields(cleanedText)
		for _, word := range words {
			wordCount[word]++
		}
		return wordCount, nil
	}
	return nil, fmt.Errorf("%w", ErrEmptyString)
}
