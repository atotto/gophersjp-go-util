package tr

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// GetDocs return translate documents filepath
func GetDocs(root string) ([]string, error) {
	pattern := regexp.MustCompilePOSIX(regexp.QuoteMeta(root) + `/(.*)(\.go|\.html)`)

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
	url string
}

func (info *revinfo) String() string {
	return info.rev
}

func (info *revinfo) URL() string {
	return info.url
}

// GetRevision returns revision string.
func GetRevision(filename string) (revinfo, error) {
	pattern := regexp.MustCompilePOSIX(`(https://code.google.com/p/go/source/browse/.*)\?.*r=(.*)$`)

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
		result := pattern.FindSubmatch(line)

		if result != nil {
			return revinfo{rev: string(result[2]), url: string(result[1])}, nil
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
