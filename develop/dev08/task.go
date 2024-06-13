package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"shell/commander"
	"strings"
)

func main() {
	readCMD()
}

// Read command from stdin and execute it
func readCMD() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			// Exit if eof signal
			if errors.Is(err, io.EOF) {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, "read input: ", err)
			continue
		}

		if err := execCMD(input); err != nil {
			fmt.Fprintln(os.Stderr, "exec cmd:", err)
			continue
		}
	}
}

// Execute command
func execCMD(input string) error {
	const pipeSign = "|"
	input = strings.TrimSpace(input)

	// If command contains pipe sign exec through pipes
	if strings.Contains(input, pipeSign) {
		if err := execPipe(input); err != nil {
			return err
		}
		return nil
	}

	cmd := strings.Fields(input)
	if len(input) == 0 {
		return nil
	}

	c := commander.NewCommander()
	switch cmd[0] {
	case "cd":
		if err := c.CD(cmd); err != nil {
			return fmt.Errorf("cd: %v", err)
		}
	case "pwd":
		dir, err := c.PWD()
		if err != nil {
			return fmt.Errorf("pwd: %v", err)
		}
		fmt.Println(dir)
	case "echo":
		c.Echo(cmd)
	case "kill":
		if err := c.Kill(cmd[1]); err != nil {
			return fmt.Errorf("kill: %v", err)
		}
	case "ps":
		if err := c.PS(); err != nil {
			return fmt.Errorf("ps: %v", err)
		}
	case "exec":
		if err := c.Exec(cmd); err != nil {
			return fmt.Errorf("exec: %v", err)
		}
	case "exit":
		os.Exit(0)
	default:
		return fmt.Errorf("command not found: %s", cmd[0])
	}
	return nil
}

// Execute piped commands
func execPipe(input string) error {
	cmds := strings.Split(input, "|")
	var lastCmdOutput bytes.Buffer

	for _, c := range cmds {
		c = strings.TrimSpace(c)
		fields := strings.Fields(c)

		if len(fields) == 0 {
			continue
		}

		cmd := exec.Command(fields[0], fields[1:]...)

		// Set input to the last command output
		cmd.Stdin = &lastCmdOutput

		var cmdOutput bytes.Buffer
		// Set output to capture command output
		cmd.Stdout = &cmdOutput

		// Run command
		err := cmd.Run()
		if err != nil {
			return err
		}
		lastCmdOutput = cmdOutput
	}
	fmt.Println(lastCmdOutput.String())
	return nil
}
