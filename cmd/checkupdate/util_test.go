package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func ExampleFind() {
	list, err := find("fixture", filePattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range list {
		fmt.Println(filepath.ToSlash(path))
	}

	// Output:
	// fixture/doc/doc.go
	// fixture/html/needupdate.html
	// fixture/html/updated.html
}

func ExampleSearch() {
	out, err := search("fixture/doc/doc.go", reposUrlPattern)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	// Output:
	// https://code.google.com/p/go/source/browse/misc/goplay/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
}

func TestIsLatest(t *testing.T) {
	ts := httptest.NewServer(http.FileServer(http.Dir("fixture/html")))
	defer ts.Close()

	tests := []struct {
		url      string
		expected bool
	}{
		{url: "/needupdate.html", expected: false},
		{url: "/updated.html", expected: true},
	}

	for n, tt := range tests {
		actual, _ := isLatest(ts.URL + tt.url)
		if actual != tt.expected {
			t.Fatalf("#%d want %v, got %v", n, tt.expected, actual)
		}
	}
}

func TestShortURL(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{
			url:      `https://code.google.com/p/go/source/browse/src/cmd/5a/doc.go?name=3633a89bb56d&amp;r=1a86e8314ff5236a92beac4a399b96efe4c407af`,
			expected: `https://code.google.com/p/go/source/browse/src/cmd/5a/doc.go?r=1a86e8314ff5236a92beac4a399b96efe4c407af`,
		},
	}

	for n, tt := range tests {
		actual := shortURL(tt.url)
		if actual != tt.expected {
			t.Fatalf("#%d got %s, want %s", n, actual, tt.expected)
		}
	}

}
