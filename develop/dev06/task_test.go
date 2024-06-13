package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	args := [][]string{
		{"go", "run", "task.go", "-d=,", "-f=2"},
		{"go", "run", "task.go", "-d=,", "-f=1,3"},
		{"go", "run", "task.go", "-d=:", "-f=2,4,6", "-s"},
	}

	input := []string{
		"apple,banana,grape",
		"apple,banana,grape",
		"Name: John, Age: 30, City: New York",
	}

	expected := []string{
		"banana",
		"apple,grape",
		"John, Age: New York",
	}

	for i, arg := range args {
		cmd := exec.Command(arg[0], arg[1:]...)

		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatalf("failed to open stdin pipe: %v", err)
		}
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		if err := cmd.Start(); err != nil {
			t.Fatalf("failed to start command: %v", err)
		}

		if _, err := io.WriteString(stdin, input[i]); err != nil {
			t.Fatalf("failed to write to stdin: %v", err)
		}
		stdin.Close()

		if err := cmd.Wait(); err != nil {
			t.Fatalf("command failed: %v", err)
		}

		if got := strings.TrimSpace(stdout.String()); got != expected[i] {
			t.Errorf("test case %d failed:\nexpected: %s\ngot: %s", i, expected[i], got)
		}
	}
}
