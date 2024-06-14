package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: wget <url>")
		os.Exit(1)
	}
	// Get URL from command line arguments
	url := os.Args[1]
	// Directory for downloaded websites
	downloadDir := "./web"

	if err := downloadSite(url, downloadDir); err != nil {
		fmt.Fprintln(os.Stderr, "failed to download site: ", err)
		os.Exit(1)
	}
}

// Downloads entire website
func downloadSite(url, dir string) error {
	// Get request to URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if response is OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download page: %s", resp.Status)
	}

	// Create directory to save downloaded files
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Extract links from the HTML webpage
	links, err := extractLinks(resp.Body)
	if err != nil {
		return err
	}

	// Iterates over each link and download its content
	for _, link := range links {
		absLink := link
		// If link is relative make it absolute
		if !strings.HasPrefix(absLink, "http") {
			absLink = url + "/" + absLink
		}

		// Get filename from link
		filename := path.Base(absLink)
		// Filepath to save downloaded file
		filepath := path.Join(dir, filename)
		if strings.HasSuffix(absLink, "/") {
			if err := downloadSite(absLink, filepath); err != nil {
				return err
			}
		} else {
			if err := downloadFile(absLink, filepath); err != nil {
				return err
			}
		}
	}

	return nil
}

// Extract links from HTML
func extractLinks(r io.Reader) ([]string, error) {
	tokenizer := html.NewTokenizer(r)
	links := make([]string, 0)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return nil, err
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}

	return links, nil
}

// Download file
func downloadFile(url, filepath string) error {
	// Make GET request to URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create file to save downloaded content
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy content of response body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("downloaded:", url)
	return nil
}
