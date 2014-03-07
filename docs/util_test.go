package docs

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
	// doc/doc2.go
	// doc/godoc.go
	// html/needupdate.html
	// html/updated.html
}

func ExampleGetTranslationTargetDocs() {
	list, err := GetTranslationTargetDocs("../_fixture")
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
	fmt.Println(out.RawURL())
	fmt.Println(out.URL())
	fmt.Println(out.String())

	// Output:
	// https://code.google.com/p/go/source/browse/misc/goplay/doc.go
	// https://code.google.com/p/go/source/browse/misc/goplay/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
	// 3633a89bb56d9276a9fe55435b849f931bfa6393
}

func ExampleGetRevision2() {
	out, err := GetRevision("../_fixture/doc/doc2.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.URL())
	fmt.Println(out.RawURL())
	fmt.Println(out.String())

	// Output:
	// https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?r=0e399fef76b7c34144d51e7b64c6da5b5591ea51&repo=tools
	// https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools
	// 0e399fef76b7c34144d51e7b64c6da5b5591ea51
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
