package main

import (
	"bufio"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

func find(root string, pattern *regexp.Regexp) ([]string, error) {
	list := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		result := pattern.MatchString(path)
		if result {
			list = append(list, path)
		}
		return nil
	})

	return list, err
}

func search(filename string, pattern *regexp.Regexp) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			return "", nil
		}
		if err != nil {
			return "", err
		}
		if isPrefix {
			return "", nil
		}
		result := pattern.Find(line)
		if result != nil {
			return string(result), nil
		}
	}
}

var nextPattern = regexp.MustCompilePOSIX(`href="(/p/go/source/browse/.*)" title="Next"`)
var host = `https://code.google.com`

func isLatest(url string) (bool, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			return true, url
		}
		if err != nil {
			log.Fatal(err)
		}
		if isPrefix {
			return true, url
		}
		results := nextPattern.FindSubmatch(line)
		if results != nil {
			newurl := host + string(results[1])
			return false, shortURL(newurl)
		}
	}

	return true, url
}

// `https://code.google.com/p/go/source/browse/src/cmd/5a/doc.go?name=3633a89bb56d&amp;r=1a86e8314ff5236a92beac4a399b96efe4c407af`
// to
// `https://code.google.com/p/go/source/browse/src/cmd/5a/doc.go?r=1a86e8314ff5236a92beac4a399b96efe4c407af`
func shortURL(uri string) string {
	uri = html.UnescapeString(uri)
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Del("name")
	u.RawQuery = q.Encode()
	return u.String()
}
