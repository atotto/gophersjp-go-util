package main

import (
	"flag"
	"log"
	"regexp"
	"time"
)

var filePattern = regexp.MustCompilePOSIX(`\.go|\.html`)
var reposUrlPattern = regexp.MustCompilePOSIX(`https://code.google.com/p/go/source/browse/.*$`)

var (
	root   = flag.String("root", ".", "root dir")
	output = flag.String("o", "", "html file for output; default: stdout")
)

//
func NewData() Data {
	d := Data{LastUpdate: time.Now().Format(time.RFC3339)}
	return d
}

type Data struct {
	Files      []*translatedFile
	LastUpdate string
}

type translatedFile struct {
	File       string
	CurrentUrl string
	NextUrl    string
	IsLatest   bool
}

func main() {
	flag.Parse()

	list, err := find(*root, filePattern)
	if err != nil {
		log.Fatal(err)
	}

	d := NewData()
	for _, path := range list {
		url, err := search(path, reposUrlPattern)
		if err != nil {
			log.Fatal(err)
		}
		isLatest, nextUrl := isLatest(url)
		d.Files = append(d.Files, &translatedFile{
			File:       path,
			CurrentUrl: url,
			NextUrl:    nextUrl,
			IsLatest:   isLatest,
		})
	}

	// Output HTML
	if *output != "" {
		err = htmlOutput(d, *output)
	}

}
