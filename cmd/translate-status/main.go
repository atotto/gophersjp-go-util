package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"sort"
	"strings"

	"go/build"

	"github.com/atotto/gophersjp-go-util/docs"
	"github.com/atotto/gophersjp-go-util/hg"
)

var (
	root   = flag.String("docroot", ".", "gophersjp/go document root dir")
	goroot = flag.String("goroot", runtime.GOROOT(), "Go root directory")
	output = flag.String("o", "translate.html", "html file for output; default: ./translate.html")
)

func main() {
	flag.Parse()

	list, err := docs.GetDocs(*root)
	if err != nil {
		log.Fatal(err)
	}

	goRepos, err := hg.AttachRepos(*goroot)
	if err != nil {
		log.Fatal(err)
	}

	doclist, err := docs.GetTranslationTargetDocs(*goroot)
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

	gotoolsDocList, err := docs.GetTranslationTargetDocs(p.Dir)
	if err != nil {
		log.Fatal(err)
	}

	num := len(doclist) + len(gotoolsDocList)
	allDocs := make(map[string]*NonTranslatedItem, num)
	for _, path := range doclist {
		allDocs[path] = &NonTranslatedItem{
			FilePath: path,
			KeyName:  path,
			RepoURL:  "https://code.google.com/p/go/source/browse/" + path,
		}
	}
	for _, path := range gotoolsDocList {
		fpath := "src/pkg/code.google.com/p/go.tools/" + path
		allDocs[fpath] = &NonTranslatedItem{
			FilePath: fpath,
			KeyName:  path,
			RepoURL:  "https://code.google.com/p/go/source/browse/" + path + "?repo=tools",
		}
	}

	d := NewDataSet()
	d.Tag = "go1.3"
	for _, path := range list {
		log.Printf("== %s\n", path)
		rev, err := docs.GetRevision(*root + "/" + path)
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		log.Printf("rev: %s\n", rev.String())
		delete(allDocs, path)

		tf := Item{
			FilePath: path,
			Rev:      rev.String(),
			RepoURL:  rev.RawURL(),
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
				tf.KeyName = rpath
				st, diff, err = gotoolRepos.Check(tag, rpath, rev.String())
				tf.Repo = "tools"
			} else {
				st, diff, err = goRepos.Check(tag, path, rev.String())
				tf.KeyName = path
				tf.Repo = "default"
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

		d.Items = append(d.Items, &tf)
	}

	// Save non-translated items
	var keys []string
	for key := range allDocs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		d.NonTranslatedItems = append(d.NonTranslatedItems, allDocs[key])
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

func createNonTranslatedDocList() map[string]bool {
	return nil
}
