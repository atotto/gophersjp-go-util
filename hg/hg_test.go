package hg

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

func TestDiff(t *testing.T) {
	//$ hg diff src/cmd/5a/doc.go -r 3633a89bb56d

	hg, err := AttachRepos(runtime.GOROOT())
	if err != nil {
		t.Fatal(err)
	}

	actual, err := hg.Diff("src/cmd/5a/doc.go", "3633a89bb56d")

	if err != nil {
		t.Fatal(err)
	}
	if len(actual) == 0 {
		t.Fatalf("want differ, got not differ")
	}
}

func TestCheck(t *testing.T) {
	//$ hg diff src/cmd/5a/doc.go -r 3633a89bb56d

	hg, err := AttachRepos(runtime.GOROOT())
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		tag            string
		path           string
		rev            string
		expectedStatus Status
		expectedDiff   int
	}{
		{"go1.2", "src/cmd/5c/doc.go", "3633a89bb56d", Same, 0},
		{"go1.2", "src/cmd/5c/doc.go", "2695fe638e0b", Outdated, 3},
		{"go1.2", "src/cmd/5a/doc.go", "3633a89bb56d", Outdated, 1},
		{"go1.2", "src/cmd/5a/doc.go", "b1edf8faa5d6", Same, 0},
		{"go1", "src/cmd/5a/doc.go", "b1edf8faa5d6", Ahead, 2},
	}

	for n, tt := range tests {
		st, diff, err := hg.Check(tt.tag, tt.path, tt.rev)
		if err != nil {
			t.Fatal(err)
		}
		if st != tt.expectedStatus {
			t.Errorf("#%d, want status=%v, got %v", n, tt.expectedStatus, st)
		}
		if diff != tt.expectedDiff {
			t.Errorf("#%d want diff=%v, got %v", n, tt.expectedDiff, diff)
		}
	}
}
