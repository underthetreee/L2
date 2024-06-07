package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"", "", errors.New("empty string")},
		{"a", "a", nil},
		{"abc", "abc", nil},
		{"3a", "", errors.New("invalid string")},
		{"2a3b", "", errors.New("invalid string")},
	}

	for _, tc := range testCases {
		result, err := unpackString(tc.input)
		if err != nil && err.Error() != tc.err.Error() {
			t.Errorf("unpackString(%q) returned error %v, expected %v", tc.input, err, tc.err)
		}
		if result != tc.expected {
			t.Errorf("unpackString(%q) returned %q, expected %q", tc.input, result, tc.expected)
		}
	}
}
