package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"greputil/printer"
	"log"
	"os"
	"regexp"
	"strings"
)

type Options struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func main() {
	// Parse command line arguments
	pattern, files, options, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// Process each file
	for _, file := range files {
		if err := processFile(file, pattern, options); err != nil {
			log.Fatal(err)
		}
	}
}

// Parses command line flags and arguments
func parseFlags() (string, []string, *Options, error) {
	// Define flags
	after := flag.Int("A", 0, "print +N lines after matching")
	before := flag.Int("B", 0, "print +N lines before matching")
	context := flag.Int("C", 0, "print +/-N lines around matching")
	count := flag.Bool("c", false, "count of matching lines")
	ignoreCase := flag.Bool("i", false, "ignore case")
	invert := flag.Bool("v", false, "invert match")
	fixed := flag.Bool("F", false, "fixed string match")
	lineNum := flag.Bool("n", false, "print line number")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return "", nil, nil, errors.New("usage: grep [flags] pattern [file1 file2 ...]")
	}

	pattern := args[0]
	files := args[1:]

	options := &Options{
		After:      *after,
		Before:     *before,
		Context:    *context,
		Count:      *count,
		IgnoreCase: *ignoreCase,
		Invert:     *invert,
		Fixed:      *fixed,
		LineNum:    *lineNum,
	}
	return pattern, files, options, nil
}

// Opens and processes each file
func processFile(file string, pattern string, options *Options) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if !options.Fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	var lines []string

	// Read lines from scanner and append to slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	printer := printer.NewPrinter(lines)

	counter := 0
	// Iterates over slice and searching match lines
	for i, line := range lines {
		matched := match(line, pattern, options)

		// Check if line matches pattern based on options
		if (options.Invert && !matched) || (!options.Invert && matched) {
			// Increases counter if flag set
			if options.Count {
				counter++
			}

			// Prints line numbers if flag set
			if options.LineNum {
				line = fmt.Sprintf("%d: %s", i+1, line)
			}

			// Prints lines around match if flag set
			if options.Context > 0 {
				printer.PrintAround(options.Context, i)
				continue
			}

			// Prints lines before match if flag set
			if options.Before > 0 {
				printer.PrintBefore(options.Before, i)

				if !printer.Found(line) {
					printer.Add(line)
					fmt.Println(line)
				}
				continue
			}

			// Prints lines after match if flag set
			if options.After > 0 {
				if !printer.Found(line) {
					printer.Add(line)
					fmt.Println(line)
				}
				printer.PrintAfter(options.After, i)
				continue
			}

			fmt.Println(line)
		}

	}

	// Prints how many lines matches if flag set
	if options.Count {
		fmt.Println(counter)
	}
	return nil
}

func match(line, pattern string, options *Options) bool {
	if options.IgnoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
		return strings.Contains(line, pattern)
	}
	return strings.Contains(line, pattern)
}
