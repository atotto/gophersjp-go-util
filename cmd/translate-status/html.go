package main

import (
	"html/template"
	"os"
	"time"
)

type Data struct {
	Files      []*translatedFile
	LastUpdate string
	Tag        string
}

func NewData() *Data {
	return &Data{LastUpdate: time.Now().Format(time.RFC3339)}
}

type translatedFile struct {
	File       string
	CurrentUrl string
	NextUrl    string
	Tip        Status
	Stable     Status
}

type Status struct {
	IsOutdated bool
	Stage      string
}

var htmlTemplate = template.Must(template.New("html").Parse(tmplHTML))

func htmlOutput(d *Data, outfile string) error {
	if outfile != "" {
		f, err := os.Create(outfile)
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(title)
		htmlTemplate.Execute(f, d)
	} else {
		htmlTemplate.Execute(os.Stdout, d)
	}

	return nil
}

const title = `<!--{
	"Title": "Translate Status"
}-->`

const tmplHTML = `
  <style>
      .status {
        margin-left: 20px;
        border-collapse: collapse;
      }
      .status td, .status th {
        vertical-align: top;
        padding: 2px 4px;
        font-size: 10pt;
      }
      .status th {
        font-size: 12pt;
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
      .timestamp {
        color: #C00;
      }
  </style>
  <h3><a href="https://code.google.com/p/go">Go</a>のドキュメント翻訳ステータス</h3>
  <p>
    翻訳したものが最新のものかどうかをチェックします。
    本サイトは github.com/gophersjp/go の成果物です。ご気軽にご参加ください。
  </p>
  <p>
  <a href="https://github.com/gophersjp/go" class="download" id="start" target="_blank">
  <span class="big">Repository</span>
  <span class="desc">github.com/gophersjp/go</span>
  </a>
  </p>
  <table class="status">
    <colgroup class="col-item"></colgroup>
    <colgroup class="col-translate"></colgroup>
    <tbody>
      <tr>
        <th><a href="https://github.com/gophersjp/go">translated item<a></th>
        <th><a href="https://code.google.com/p/go/source/browse">{{.Tag}}</a></th>
        <th><a href="https://code.google.com/p/go/source/browse">tip</a></th>
      </tr>
      {{range $i, $f := .Files}}
      <tr class="item">
        <td><a href="https://github.com/gophersjp/go/blob/master/{{$f.File}}">{{$f.File}}</a></td>
        <td>
          {{if $f.Stable.IsOutdated}}
            <a href="{{$f.NextUrl}}" class="outdated">{{$f.Stable.Stage}}(next)</a>
          {{else}}
            <a href="{{$f.NextUrl}}" class="latest">{{$f.Stable.Stage}}</a>
          {{end}}
        </td>
        <td>
          {{if $f.Tip.IsOutdated}}
            <a href="{{$f.NextUrl}}" class="latest">{{$f.Tip.Stage}}</a>
          {{else}}
            <a href="{{$f.NextUrl}}" class="latest">{{$f.Tip.Stage}}</a>
          {{end}}
        </td>
      </tr>
      {{end}}
    </tbody>
  </table>
  <div id="footer">
    Last Update: {{.LastUpdate}}
  </div>
`
