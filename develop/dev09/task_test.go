package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestExtractLinks(t *testing.T) {
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Test Page</title>
	</head>
	<body>
	<a href="https://example.com">Link 1</a>
	<a href="/page2">Link 2</a>
	<a href="http://example.org">Link 3</a>
	</body>
	</html>
	`

	expectedLinks := []string{"https://example.com", "/page2", "http://example.org"}

	r := strings.NewReader(htmlContent)
	links, err := extractLinks(r)
	if err != nil {
		t.Fatalf("extractLinks error: %v", err)
	}

	if len(links) != len(expectedLinks) {
		t.Fatalf("expected %d links, got %d", len(expectedLinks), len(links))
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("expected link '%s', got '%s'", expectedLinks[i], link)
		}
	}
}

func TestDownloadFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test content"))
	}))
	defer server.Close()

	tmpFile, err := os.CreateTemp("", "test_download_file")
	if err != nil {
		t.Fatalf("TempFile error: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = downloadFile(server.URL, tmpFile.Name())
	if err != nil {
		t.Fatalf("downloadFile error: %v", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	expectedContent := []byte("test content")
	if !bytes.Equal(content, expectedContent) {
		t.Errorf("expected content '%s', got '%s'", expectedContent, content)
	}
}
