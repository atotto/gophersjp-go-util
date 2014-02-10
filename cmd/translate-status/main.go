package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"

	"go/build"
	"github.com/atotto/gophersjp-go-util/hg"
	"github.com/atotto/gophersjp-go-util/translated"
)

var (
	root   = flag.String("docroot", ".", "gophersjp/go document root dir")
	goroot = flag.String("goroot", runtime.GOROOT(), "Go root directory")
	output = flag.String("o", "translate.html", "html file for output; default: ./translate.html")
)

func main() {
	flag.Parse()

	list, err := tr.GetDocs(*root)
	if err != nil {
		log.Fatal(err)
	}

	goRepos, err := hg.AttachRepos(*goroot)
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
	d.Tag = "go1.2"
	for _, path := range list {
		log.Printf("== %s\n", path)

		rev, err := tr.GetRevision(*root + "/" + path)
		if err != nil {
			log.Fatal("rev:" + err.Error())
		}
		tf := translatedFile{
			File:       path,
			CurrentUrl: rev.String(),
			NextUrl:    rev.RawURL(),
			Stable: Status{
				IsOutdated: false,
				Stage:      "",
			},
			Tip: Status{
				IsOutdated: false,
				Stage:      "",
			},
		}

		fn := func(s *Status, tag string) {
			var st hg.Status
			var diff int
			if strings.HasPrefix(path, "src/pkg/code.google.com/p/go.tools/") {
				rpath := strings.TrimPrefix(path, "src/pkg/code.google.com/p/go.tools/")
				st, diff, err = gotoolRepos.Check(tag, rpath, rev.String())
			} else {
				st, diff, err = goRepos.Check(tag, path, rev.String())
			}
			switch st {
			case hg.Same:
				s.IsOutdated = false
				s.Stage = "OK"
			case hg.Ahead:
				s.IsOutdated = false
				s.Stage = fmt.Sprintf("ahead+%d", diff)
			case hg.Outdated:
				s.IsOutdated = true
				s.Stage = fmt.Sprintf("outdated-%d", diff)
			default:
				s.IsOutdated = true
				s.Stage = "error"
			}
			if err != nil {
				s.IsOutdated = false
				log.Printf("error: %s", err.Error())
				s.Stage = "none"
			}
		}
		fn(&tf.Stable, d.Tag)
		fn(&tf.Tip, "tip")

		d.Files = append(d.Files, &tf)
	}

	// Output HTML
	err = htmlOutput(d, *output)
	log.Printf("Output: %s\n", *output)
	if err != nil {
		log.Fatal(err)
	}
}

func check(tag, filepath, revision string) {

}
