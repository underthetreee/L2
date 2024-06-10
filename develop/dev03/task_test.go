package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadWords(t *testing.T) {
	filename := "testfile.txt"
	content := "word1\nword2\nword3"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	words, err := readWords(filename)

	if err != nil {
		t.Errorf("readWords(%q) returned error: %v", filename, err)
	}

	expected := []string{"word1", "word2", "word3"}
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("readWords(%q) = %v, want %v", filename, words, expected)
	}
}

func TestWriteWords(t *testing.T) {
	words := []string{"word1", "word2", "word3"}
	filename := "testfile.txt"

	err := writeWords(filename, words)
	if err != nil {
		t.Errorf("writeWords(%q, %v) returned error: %v", filename, words, err)
	}

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	expectedContent := "word1\nword2\nword3\n"
	if string(fileContent) != expectedContent {
		t.Errorf(
			"writeWords(%q, %v) produced incorrect content: got %q, want %q",
			filename,
			words,
			string(fileContent),
			expectedContent,
		)
	}
}

func TestGetColumn(t *testing.T) {
	word := "apple banana cherry"
	tests := []struct {
		index    int
		expected string
	}{
		{1, "apple"},
		{2, "banana"},
		{3, "cherry"},
		{4, ""},
	}

	for _, test := range tests {
		result := getColumn(word, test.index)
		if result != test.expected {
			t.Errorf("getColumn(%q, %d) = %q, want %q", word, test.index, result, test.expected)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "apple", "banana"}
	expected := []string{"apple", "banana", "cherry"}
	result := removeDuplicates(words)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeDuplicates(%v) = %v, want %v", words, result, expected)
	}
}
