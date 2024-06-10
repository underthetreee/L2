package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Define command-line flags
	columnFlag := flag.Int("k", 0, "column sort")
	numericalFlag := flag.Bool("n", false, "numeric sort")
	reverseFlag := flag.Bool("r", false, "reverse sort")
	uniqueFlag := flag.Bool("u", false, "unique sort")
	flag.Parse()

	// Get filename from command-line arguments
	filename := flag.Arg(0)
	// Reads words from file
	words, err := readWords(filename)
	if err != nil {
		log.Fatal(err)
	}

	// If no flags provided, perform default sorting
	if *columnFlag == 0 && !*numericalFlag && !*reverseFlag && !*uniqueFlag {
		sort.Strings(words)
	}

	// Remove dublicates if uniqueFlag set
	if *uniqueFlag {
		words = removeDuplicates(words)
	}

	// Sort words based on column if columnFlag set
	if *columnFlag > 0 {
		sort.Slice(words, func(i, j int) bool {
			col1 := getColumn(words[i], *columnFlag)
			col2 := getColumn(words[j], *columnFlag)
			if *numericalFlag {
				num1, _ := strconv.Atoi(col1)
				num2, _ := strconv.Atoi(col2)
				return num1 < num2
			}
			return col1 < col2
		})
	}

	// Reverse order if reverseFlag set
	if *reverseFlag {
		sort.Sort(sort.Reverse(sort.StringSlice(words)))
	}

	// Write sorted words
	if err := writeWords(filename, words); err != nil {
		log.Fatalln(err)
	}
}

// Reads words from file and returns slice of strings
func readWords(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w := scanner.Text()
		words = append(words, w)
	}
	return words, nil
}

// Writes words to a file
func writeWords(filename string, words []string) error {
	f, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, w := range words {
		_, err := writer.WriteString(w + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

// Extracts a specific column from a space-separated string
func getColumn(word string, columnIndex int) string {
	fields := strings.Fields(word)
	if columnIndex <= len(fields) {
		return fields[columnIndex-1]
	}
	return ""
}

// Removes duplicate words from a slice
func removeDuplicates(words []string) []string {
	uniqueWords := make(map[string]bool)
	result := make([]string, 0)

	for _, w := range words {
		if !uniqueWords[w] {
			uniqueWords[w] = true
			result = append(result, w)
		}
	}
	return result
}
