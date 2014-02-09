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
	output = flag.String("o", "translate.html", "html file for output; default: stdout")
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
			if strings.HasPrefix(path, "src/pkg/code.google.com/p/go.tools/") {
				rpath := strings.TrimPrefix(path, "src/pkg/code.google.com/p/go.tools/")
				st, _, err = gotoolRepos.Check(tag, rpath, rev.String())
			} else {
				st, _, err = goRepos.Check(tag, path, rev.String())
			}
			switch st {
			case hg.Same:
				s.IsOutdated = false
				s.Stage = "OK"
			case hg.Ahead:
				s.IsOutdated = false
				s.Stage = "Ahead"
			case hg.Outdated:
				s.IsOutdated = true
				s.Stage = "Outdated"
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
