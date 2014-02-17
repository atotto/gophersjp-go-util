package main

func ExampleHtmlOutput() {
	d := NewData()
	d.Tag = "go1.2"
	d.Files = append(d.Files, &translatedFile{
		File:       "tools/cmd/godoc/doc.go",
		KeyName:    "cmd/godoc/doc.go",
		CurrentUrl: "https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools&r=0e399fef76b7c34144d51e7b64c6da5b5591ea51",
		RepoURL:    "https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools",
		Tip: Status{
			IsOutdated: false,
			Stage:      "hoge",
		},
		Stable: Status{
			IsOutdated: false,
			Stage:      "hoge",
		},
	})
	d.Files = append(d.Files, &translatedFile{
		File:       "src/cmd/gofmt/doc.go",
		CurrentUrl: "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
		RepoURL:    "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
	})

	htmlOutput(d, "update.html")

	// Output:
}
