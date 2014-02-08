package main

import (
	"flag"
	"log"
	"runtime"
	"strings"

	"go/build"
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

	goRepos, err := hg.AttachRepos(runtime.GOROOT())
	if err != nil {
		log.Fatal(err)
	}

	p, err := build.Import("code.google.com/p/go.tools", build.Default.GOROOT, build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}

	gotoolRepos, err := hg.AttachRepos(p.Dir)
	if err != nil {
		log.Fatal(err)
	}

	d := NewData()
	for _, path := range list {
		log.Printf("== %s\n", path)

		rev, err := tr.GetRevision(*root + "/" + path)
		if err != nil {
			log.Fatal("rev:" + err.Error())
		}
		var b []byte
		if strings.HasPrefix(path, "src/pkg/code.google.com/p/go.tools") {
			b, err = gotoolRepos.Diff(path, rev.String())
		} else {
			b, err = goRepos.Diff(path, rev.String())
		}
		var revision = "LATEST"
		if len(b) != 0 {
			revision = rev.String()
		}
		if err != nil {
			log.Printf("error: %s, %s", err.Error(), string(b))
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
