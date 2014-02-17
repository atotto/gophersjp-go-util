package main

func ExampleHtmlOutput() {
	d := NewDataSet()
	d.Tag = "go1.2"
	d.Items = append(d.Items, &Item{
		FilePath: "tools/cmd/godoc/doc.go",
		KeyName:  "cmd/godoc/doc.go",
		Rev:      "https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools&r=0e399fef76b7c34144d51e7b64c6da5b5591ea51",
		RepoURL:  "https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools",
		Tip: Status{
			IsOutdated: false,
			Stage:      "hoge",
		},
		Stable: Status{
			IsOutdated: false,
			Stage:      "hoge",
		},
	})
	d.Items = append(d.Items, &Item{
		FilePath: "src/cmd/gofmt/doc.go",
		Rev:      "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
		RepoURL:  "https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a",
	})

	htmlOutput(d, "update.html")

	// Output:
}
