package main

import (
	"os"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		line     string
		pattern  string
		options  *Options
		expected bool
	}{
		{
			line:     "This is a test line.",
			pattern:  "test",
			options:  &Options{IgnoreCase: false},
			expected: true,
		},
		{
			line:     "This is a test line.",
			pattern:  "TEST",
			options:  &Options{IgnoreCase: true},
			expected: true,
		},
	}

	for _, test := range tests {
		result := match(test.line, test.pattern, test.options)
		if result != test.expected {
			t.Errorf("For line: %s, pattern: %s, expected: %v, got: %v", test.line, test.pattern, test.expected, result)
		}
	}
}

func TestProcessFile(t *testing.T) {
	tempFile := createTempFile(t, "line 1\nline 2\nline 3\n")

	options := &Options{IgnoreCase: false}
	err := processFile(tempFile, "line 2", options)
	if err != nil {
		t.Errorf("Error processing file: %v", err)
	}

	options = &Options{IgnoreCase: true}
	err = processFile(tempFile, "LINE 3", options)
	if err != nil {
		t.Errorf("Error processing file: %v", err)
	}

	options = &Options{IgnoreCase: false}
	err = processFile(tempFile, "not found", options)
	if err != nil {
		t.Errorf("Error processing file: %v", err)
	}
}

func createTempFile(t *testing.T, content string) string {
	tempFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer tempFile.Close()

	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	return tempFile.Name()
}
