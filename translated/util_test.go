package tr

import (
	"fmt"
	"log"
	"path/filepath"
)

func ExampleGetDocs() {
	list, err := GetDocs("../fixture")
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range list {
		fmt.Println(filepath.ToSlash(path))
	}

	// Output:
	// ../fixture/doc/doc.go
	// ../fixture/html/needupdate.html
	// ../fixture/html/updated.html
}

func ExampleGetNode() {
	out, err := GetNode("../fixture/doc/doc.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	// Output:
	// https://code.google.com/p/go/source/browse/misc/goplay/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
}
