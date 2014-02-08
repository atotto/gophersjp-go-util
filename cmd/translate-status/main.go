package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/atotto/gophersjp-go-util/hg"
	"github.com/atotto/gophersjp-go-util/translated"
)

var (
	root   = flag.String("root", ".", "root dir")
	output = flag.String("o", "update.html", "html file for output; default: stdout")
)

func main() {
	flag.Parse()

	list, err := tr.GetDocs(*root)
	if err != nil {
		log.Fatal(err)
	}

	repos, err := hg.AttachRepos(runtime.GOROOT())
	if err != nil {
		log.Fatal(err)
	}

	d := NewData()
	for _, path := range list {
		rev, err := tr.GetRevision(*root + "/" + path)
		if err != nil {
			log.Fatal("rev:" + err.Error())
		}
		b, err := repos.Diff(path, rev.String())
		var revision = "LATEST"
		if len(b) != 0 {
			revision = rev.String()
		}
		if err != nil {
			log.Printf("path: %s, %s, %s", path, err.Error(), string(b))
			revision = string(b)
		}

		d.Files = append(d.Files, &translatedFile{
			File:       path,
			CurrentUrl: rev.String(),
			NextUrl:    rev.URL(),
			Revision:   revision,
		})
	}

	// Output HTML
	err = htmlOutput(d, *output)
	log.Printf("Output: %s\n", *output)
	if err != nil {
		log.Fatal(err)
	}
}
