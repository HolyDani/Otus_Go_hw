package main

import (
	"os"
	"testing"
)

func TestOpenRead(t *testing.T) {
	strExpected := "LOG ERROR 1\nLOG WARNING 1"
	data, err := openRead("./test_data/testfile.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if string(data) != strExpected {
		t.Errorf("expected: %s, got: %s", strExpected, data)
	}
}

func TestOpenWrite(t *testing.T) {
	data := []string{"LOG ERROR 1\n", "LOG WARNING 1"}
	outputPath = "./test_data/testoutput.txt"
	logLevel = "WARNING"
	err := openWrite(data)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	outputData, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	expectedOutput := "LOG WARNING 1\n"
	if string(outputData) != expectedOutput {
		t.Errorf("Expected output %s, got %s", expectedOutput, outputData)
	}
}
