package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func getNTPTime(server string) (time.Time, error) {
	// Resolve NTP server address
	ips, err := net.LookupIP(server)
	if err != nil {
		return time.Time{}, err
	}

	// Get time from NTP server
	t, err := ntp.Time(ips[0].String())
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func main() {
	t, err := getNTPTime("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(t.Format(time.RFC3339))
}
