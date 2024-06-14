package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	// Check if host and port provided
	if flag.NArg() != 2 {
		fmt.Println("usage: telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	// Connect to server with temeout
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error connecting to server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create channel to wait for interrupt signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	// Read data from socket and print to stdout
	go func() {
		for {
			data := make([]byte, 4096)
			n, err := conn.Read(data)
			if err != nil {
				fmt.Fprintln(os.Stderr, "connection closed by server")
				os.Exit(0)
			}
			fmt.Print(string(data[:n]))
		}
	}()

	// Read input data and write it to socket
	go func() {
		for {
			data := make([]byte, 4096)
			n, err := os.Stdin.Read(data)
			if err != nil {
				fmt.Fprintln(os.Stderr, "reading from stdin: ", err)
				os.Exit(1)
			}
			if n == 0 {
				continue
			}
			_, err = conn.Write(data[:n])
			if err != nil {
				fmt.Fprintln(os.Stderr, "writing to server: ", err)
				os.Exit(1)
			}
		}
	}()

	<-done
}
