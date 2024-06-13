package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	Fields    string
	Delimiter string
	Separated bool
}

func main() {
	options := parseFlags()

	if err := processInput(options); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Reads input from stdin and processes it based on provided options
func processInput(options *Options) error {
	scanner := bufio.NewScanner(os.Stdin)

	// Iterate over each line in the input
	for scanner.Scan() {
		line := scanner.Text()

		// If separated flag is set and line doesn't contain delimiter, skip it
		if options.Separated && !strings.Contains(line, options.Delimiter) {
			continue
		}

		// Split line into fields using specified delimiter
		fields := strings.Split(line, options.Delimiter)

		// If specific fields are selected, filter and print only those fields
		if options.Fields != "" {
			var selectedFields []string
			for _, f := range strings.Split(options.Fields, ",") {
				idx := parseFieldIndex(f)
				// Ensure index in range of available fields
				if idx >= 0 && idx < len(fields) {
					selectedFields = append(selectedFields, fields[idx])
				}
			}
			// Print selected fields joined by the delimiter
			fmt.Println(strings.Join(selectedFields, options.Delimiter))
		} else {
			fmt.Println(line)
		}
	}
	return nil
}

// Parses the command line flags and returns the Options struct
func parseFlags() *Options {
	options := &Options{}

	// Parse flags
	flag.StringVar(&options.Fields, "f", "", "select only these fields")
	flag.StringVar(&options.Delimiter, "d", "\t", "set delimiter")
	flag.BoolVar(&options.Separated, "s", false, "do not print lines containing delimiters")
	flag.Parse()

	return options
}

// Parses the field index from string and returns it
func parseFieldIndex(field string) int {
	var index int
	fmt.Sscanf(field, "%d", &index)
	return index - 1
}
