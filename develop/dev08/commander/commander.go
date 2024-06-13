package commander

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

type Commander struct{}

func NewCommander() *Commander {
	return &Commander{}
}

// Kill process by pid
func (c *Commander) Kill(processID string) error {
	pid, err := strconv.Atoi(processID)
	if err != nil {
		return err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	if err := process.Kill(); err != nil {
		return err
	}
	return nil
}

// Print processes
func (c *Commander) PS() error {
	procs, err := ps.Processes()
	if err != nil {
		return err
	}
	fmt.Println("PID", "Executable")
	for _, p := range procs {
		fmt.Println(p.Pid(), p.Executable())
	}
	return nil
}

// Execute provided process
func (c *Commander) Exec(cmd []string) error {
	if len(cmd) < 2 {
		return nil
	}
	bin, err := exec.LookPath(cmd[1])
	if err != nil {
		return err
	}
	env := os.Environ()
	return syscall.Exec(bin, cmd[1:], env)
}

// Change directory
func (c *Commander) CD(cmd []string) error {
	if len(cmd) < 2 {
		if err := homeDir(); err != nil {
			return err
		}
		return nil
	}
	if err := os.Chdir(cmd[1]); err != nil {
		return err
	}
	return nil
}

// Return current directory
func (c *Commander) PWD() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Print command
func (c *Commander) Echo(cmd []string) {
	str := strings.Join(cmd[1:], " ")
	fmt.Println(str)
}

// Set user home dir as current directory
func homeDir() error {
	homeDir, err := os.UserHomeDir()
	os.Chdir(homeDir)
	if err != nil {
		return err
	}
	return nil
}
