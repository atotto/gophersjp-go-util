package docs

import (
	"bufio"
	"html"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

// GetDocs return translated documents from filepath
func GetDocs(root string) ([]string, error) {
	pattern := regexp.MustCompilePOSIX(regexp.QuoteMeta(root) + `/(.*)(\.go|\.html)$`)

	list := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		result := pattern.FindStringSubmatch(path)
		if result != nil {
			list = append(list, result[1]+result[2])
		}
		return nil
	})

	return list, err
}

// GetTranslationTargetDocs return target of translation documents from filepath
func GetTranslationTargetDocs(root string) ([]string, error) {
	pattern := regexp.MustCompilePOSIX(regexp.QuoteMeta(root) + `/(.*)(doc.go|\.html)$`)

	list := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		result := pattern.FindStringSubmatch(path)
		if result != nil {
			list = append(list, result[1]+result[2])
		}
		return nil
	})

	return list, err
}

type revinfo struct {
	rev string
	url *url.URL
}

func (info *revinfo) String() string {
	return info.rev
}

func (info *revinfo) RawURL() string {
	u := *info.url // copy
	q := u.Query()
	q.Del("r")
	u.RawQuery = q.Encode()
	return u.String()
}

func (info *revinfo) URL() string {
	u := info.url
	return u.String()
}

// GetRevision returns revision string.
func GetRevision(filename string) (revinfo, error) {
	pattern := regexp.MustCompilePOSIX(`https://code.google.com/p/go/source/browse/.*$`)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			return revinfo{}, nil
		}
		if err != nil {
			return revinfo{}, err
		}
		if isPrefix {
			return revinfo{}, nil
		}
		result := pattern.Find(line)

		if result != nil {
			uri := html.UnescapeString(string(result))
			u, err := url.Parse(uri)
			if err != nil {
				log.Fatal(err)
			}
			q := u.Query()
			q.Del("name")
			u.RawQuery = q.Encode()

			return revinfo{rev: q.Get("r"), url: u}, nil
		}
	}
}

func GetURL(filename string) (string, error) {
	pattern := regexp.MustCompilePOSIX(`https://code.google.com/p/go/source/browse/.*$`)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			return "", nil
		}
		if err != nil {
			return "", err
		}
		if isPrefix {
			return "", nil
		}
		result := pattern.Find(line)
		if result != nil {
			return string(result), nil
		}
	}
}
