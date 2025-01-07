package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var filePath, outputPath, logLevel string

func main() {
	flag.StringVar(&filePath, "file", os.Getenv("LOG_ANALYZER_FILE"), "Path to log file")
	flag.StringVar(&outputPath, "output", os.Getenv("LOG_ANALYZER_OUTPUT"), "Path to output file")
	flag.StringVar(&logLevel, "level", os.Getenv("LOG_ANALYZER_LEVEL"), "")
	flag.Parse()

	data, err := openRead(filePath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	lines := strings.Split(string(data), "\n")
	if outputPath != "" {
		if err = openWrite(lines); err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		for _, line := range lines {
			if strings.Contains(line, logLevel) {
				fmt.Println(line)
			}
		}
	}
}

func openRead(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("openRead() unable to open file: %w", err)
	}
	defer file.Close()

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("openRead(): ReadFile(): unable to read all file: %w", err)
	}
	return data, nil
}

func openWrite(text []string) error {
	file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("openWrite() unable to open file: %w", err)
	}
	defer file.Close()

	for _, line := range text {
		if strings.Contains(line, logLevel) {
			_, err = file.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("openWrite(): unable to write data to file: %w", err)
			}
		}
	}
	return nil
}
