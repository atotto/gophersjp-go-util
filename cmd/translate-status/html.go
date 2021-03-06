package main

import (
	"html/template"
	"os"
	"time"
)

type DataSet struct {
	Items              []*Item
	NonTranslatedItems []*NonTranslatedItem
	LastUpdate         string
	Tag                string
}

func NewDataSet() *DataSet {
	return &DataSet{LastUpdate: time.Now().Format(time.RFC3339)}
}

type Item struct {
	FilePath string
	KeyName  string
	Rev      string
	RepoURL  string
	Repo     string
	Tip      Status
	Stable   Status
}

type Status struct {
	IsOutdated bool
	Stage      string
}

type NonTranslatedItem struct {
	FilePath string
	KeyName  string
	RepoURL  string
}

var htmlTemplate = template.Must(template.New("html").Parse(tmplHTML))

func htmlOutput(d *DataSet, outfile string) error {
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
      .col-translate {
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
      a.buttons {
        padding: 0px 5px;
        color: #222;
        border: 1px solid #375EAB;
        background: #E0EBF5;

        -webkit-border-radius: 4px;
        -moz-border-radius: 4px;
        border-radius: 4px;
      }
  </style>
  <h3><a href="https://code.google.com/p/go" target="_blank">Go</a>翻訳ステータス</h3>
  <p>
    翻訳したドキュメントが最新のものかどうかをチェックしています。
    なお、本サイトは github.com/gophersjp/go の成果物です。ご気軽にご参加ください。
    翻訳活動に参加していただける方は、<a href="https://github.com/gophersjp/go/wiki/%E7%BF%BB%E8%A8%B3%E5%8D%94%E5%8A%9B%E3%83%95%E3%83%AD%E3%83%BC" target="_blank">翻訳協力フロー</a>を一読の後、<a href="#non_translated">未翻訳リスト</a>からissueを登録ください。
  </p>
  <p>
  <a href="https://github.com/gophersjp/go" class="download" id="start" target="_blank">
  <span class="big">Repository</span>
  <span class="desc">github.com/gophersjp/go</span>
  </a>
  </p>
  <table class="status">
    <colgroup class="col-item"></colgroup>
    <colgroup class="col-translate" span="2"></colgroup>
    <colgroup class="col-issues"></colgroup>
    <tbody>
      <tr>
        <th><a href="https://github.com/gophersjp/go" target="_blank">translated item<a></th>
        <th><a href="https://code.google.com/p/go/source/browse" target="_blank">{{.Tag}}</a></th>
        <th><a href="https://code.google.com/p/go/source/browse" target="_blank">tip</a></th>
        <th colspan="2"><a href="https://github.com/gophersjp/go/issues?state=open" target="_blank">issues</a></th>
      </tr>
      {{range $i, $f := .Items}}
      <tr class="item">
        <td><a href="https://github.com/gophersjp/go/blob/master/{{$f.FilePath}}">{{$f.FilePath}}</a></td>
        <td>
          {{if $f.Stable.IsOutdated}}
            <a href="{{$f.RepoURL}}?r={{$f.Rev}}" class="outdated" target="_blank">{{$f.Stable.Stage}}</a>
          {{else}}
            <a href="{{$f.RepoURL}}" class="latest" target="_blank">{{$f.Stable.Stage}}</a>
          {{end}}
        </td>
        <td>
          {{if $f.Tip.IsOutdated}}
            <a href="{{$f.RepoURL}}?r={{$f.Rev}}" class="latest" target="_blank">{{$f.Tip.Stage}}</a>
            <a href="https://code.google.com/p/go/source/list?path=/{{$f.KeyName}}" target="_blank">history</a>
          {{else}}
            <a href="{{$f.RepoURL}}" class="latest" target="_blank">{{$f.Tip.Stage}}</a>
          {{end}}
        </td>
        <td>
          <a href="https://github.com/gophersjp/go/search?type=Issues&q={{$f.KeyName}}" target="_blank">issues</a>
        </td>
        <td>
          <a href="https://github.com/gophersjp/go/issues/new?labels=translation&title=%2e%2f{{$f.KeyName}}&body=link%3a%20{{$f.RepoURL}}" target="_blank">add issue</a>
        </td>
      </tr>
      {{end}}
    </tbody>
  </table>
  <h4 id="non_translated">未翻訳リスト</h4>
  <table class="status">
    <colgroup class="col-item"></colgroup>
    <colgroup class="col-translate" span="2"></colgroup>
    <colgroup class="col-issues"></colgroup>
    <tbody>
      <tr>
        <th><a href="https://github.com/gophersjp/go" target="_blank">non-translated item<a></th>
        <th colspan="2">翻訳用チケット</a></th>
      </tr>
      {{range $i, $f := .NonTranslatedItems}}
      <tr class="item">
        <td><a href="{{$f.RepoURL}}" target="_blank">{{$f.FilePath}}</a></td>
        <td>
          <a href="https://github.com/gophersjp/go/search?type=Issues&q={{$f.KeyName}}" target="_blank">検索</a>
        </td>
        <td>
          <a href="https://github.com/gophersjp/go/issues/new?labels=translation&title=%2e%2f{{$f.KeyName}}&body=link%3a%20{{$f.RepoURL}}" target="_blank">翻訳します！</a>
        </td>
      </tr>
      {{end}}
    </tbody>
  </table>

  <div id="footer">
    Last Update: {{.LastUpdate}}
  </div>
`
