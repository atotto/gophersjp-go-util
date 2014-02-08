package tr

import (
	"fmt"
	"log"
	"path/filepath"
)

func ExampleGetDocs() {
	list, err := GetDocs("../_fixture")
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range list {
		fmt.Println(filepath.ToSlash(path))
	}

	// Output:
	// doc/doc.go
	// html/needupdate.html
	// html/updated.html
}

func ExampleGetRevision() {
	out, err := GetRevision("../_fixture/doc/doc.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.URL())
	fmt.Println(out.String())

	// Output:
	// https://code.google.com/p/go/source/browse/misc/goplay/doc.go
	// 3633a89bb56d9276a9fe55435b849f931bfa6393
}

func ExampleGetURL() {
	out, err := GetURL("../_fixture/doc/doc.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	// Output:
	// https://code.google.com/p/go/source/browse/misc/goplay/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
}
