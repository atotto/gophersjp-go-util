package main

func ExampleHtmlOutput() {
	d := NewData()
	d.Files = append(d.Files, &translatedFile{
		File:       "tools/cmd/godoc/doc.go",
		CurrentUrl: "https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools&r=0e399fef76b7c34144d51e7b64c6da5b5591ea51",
		Revision:   "LATEST",
	})
	d.Files = append(d.Files, &translatedFile{
		File:       "src/cmd/gofmt/doc.go",
		CurrentUrl: "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
		NextUrl:    "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
		Revision:   "test",
	})
	d.Files = append(d.Files, &translatedFile{
		File:       "tools/cmd/godoc/doc.go",
		CurrentUrl: "https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools&r=0e399fef76b7c34144d51e7b64c6da5b5591ea51",
		Revision:   "LATEST",
	})
	d.Files = append(d.Files, &translatedFile{
		File:       "src/cmd/gofmt/doc.go",
		CurrentUrl: "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
		NextUrl:    "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
		Revision:   "test",
	})

	htmlOutput(d, "update.html")

	// Output:
}
