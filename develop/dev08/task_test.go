package main

import (
	"testing"
)

func TestExecCMD(t *testing.T) {
	tests := []struct {
		input        string
		expectedErr  bool
		expectedExit bool
	}{
		{input: "pwd", expectedErr: false, expectedExit: false},
		{input: "echo Hello, world!", expectedErr: false, expectedExit: false},
		{input: "cd /tmp", expectedErr: false, expectedExit: false},
		{input: "ps", expectedErr: false, expectedExit: false},
		{input: "kill 123", expectedErr: false, expectedExit: false},
		{input: "exec ls -l", expectedErr: false, expectedExit: false},
		{input: "exit", expectedErr: false, expectedExit: true},
	}

	for _, test := range tests {
		t.Logf("Running test with input: %s", test.input)
		err := execCMD(test.input)
		if test.expectedErr && err == nil {
			t.Errorf("Expected an error, but got nil")
		}
		if !test.expectedErr && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if test.expectedExit {
			break
		}
	}
}
