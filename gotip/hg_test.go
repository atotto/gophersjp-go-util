package gotip

import (
	"runtime"
	"testing"
)

func TestAttachRepos(t *testing.T) {
	{
		_, err := AttachRepos(runtime.GOROOT() + "hoge")
		if err == nil {
			t.Errorf("want no such file or directory, got nil")
		}
	}

	{
		_, err := AttachRepos(runtime.GOROOT() + "/src")
		if err == nil {
			t.Errorf("want no such file or directory, got nil")
		}
	}

	{
		_, err := AttachRepos(runtime.GOROOT())
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIsLatest(t *testing.T) {
	//$ hg diff src/cmd/5a/doc.go -r 3633a89bb56d

	hg, err := AttachRepos(runtime.GOROOT())
	if err != nil {
		t.Fatal(err)
	}

	expected := false
	actual, err := hg.IsLatest("src/cmd/5a/doc.go", "3633a89bb56d")

	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("want %v, got %v", expected, actual)
	}
}
