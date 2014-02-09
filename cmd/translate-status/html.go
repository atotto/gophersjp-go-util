package main

import (
	"html/template"
	"os"
	"time"
)

type Data struct {
	Files      []*translatedFile
	LastUpdate string
}

func NewData() *Data {
	return &Data{LastUpdate: time.Now().Format(time.RFC3339)}
}

type translatedFile struct {
	File       string
	CurrentUrl string
	NextUrl    string
	IsLatest   bool
	Revision   string
}

var htmlTemplate = template.Must(template.New("html").Parse(tmplHTML))

func htmlOutput(d *Data, outfile string) error {
	if outfile != "" {
		f, err := os.Create(outfile)
		if err != nil {
			return err
		}
		defer f.Close()
		htmlTemplate.Execute(f, d)
	} else {
		htmlTemplate.Execute(os.Stdout, d)
	}

	return nil
}

const tmplHTML = `
<!DOCTYPE html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>Gophersjp Translate Status</title>
  <style>
      body {
        font-family: sans-serif;
        padding: 0; margin: 0;
      }
      h1, h2 {
        margin: 0;
        padding: 5px;
      }
      h1 {
        background: #eee;
      }
      h2 {
        margin-top: 20px;
      }
      #content {
        margin: 10px;
      }
      .status, .packages {
        margin: 5px;
        border-collapse: collapse;
      }
      .status td, .status th, .packages td, .packages th {
        vertical-align: top;
        padding: 2px 4px;
        font-size: 10pt;
      }
      .status tr.item:nth-child(2n) {
        background-color: #f0f0f0;
      }
      .status .hash {
        font-family: monospace;
        font-size: 9pt;
      }
      .status .result {
        text-align: center;
        width: 2em;
      }
      .col-item {
        border-right: solid 1px #ccc;
      }
      .status .latest {
        color: #000;
        font-size: 83%;
      }
      .outdated {
        color: #C00;
      }
      #footer {
        clear: both;
        border-top: 1px solid #bbb;
        font-size: 0.9em;
        color: #aaa;
        padding: 5px;
        background: #fff;
      }
  </style>
</head>  
<body>
  <h1>Translate Status</h1>
  <div id="content">
    <p><a href="https://code.google.com/p/go">Go</a>のドキュメント翻訳ステータス</p>
    <p>Repository: <a href="https://github.com/gophersjp/go">github.com/gophersjp/go</a></p>
    <table class="status">
      <colgroup class="col-item"></colgroup>
      <colgroup class="col-translate"></colgroup>
      <tbody>
        <tr>
          <th>item</th>
          <th>translate revision</th>
        </tr>
        {{range $i, $f := .Files}}
        <tr class="item">
          <td><a href="https://github.com/gophersjp/go/blob/master/{{$f.File}}">{{$f.File}}</a></td>
          <td>
            {{if $f.IsLatest}}
              <a href="{{$f.NextUrl}}" class="latest">{{$f.Revision}}</a>
            {{else}}
              <span class="outdated">outdated</span>
              <a href="{{$f.NextUrl}}" class="outdated">(next)</a>
            {{end}}
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>
  <div id="footer">
    Last Update: {{.LastUpdate}}
  </div>
</div>
</body>
</html>
`
