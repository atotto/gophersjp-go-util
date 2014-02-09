package hg

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
)

type Repos struct {
	repoRoot string
}

// AttachRepos is handle of hg command.
// repoRoot is represent hg repository root.
func AttachRepos(repoRoot string) (*Repos, error) {
	if _, err := os.Stat(repoRoot); err != nil {
		return nil, err
	}
	if _, err := os.Stat(repoRoot + "/.hg"); err != nil {
		return nil, err
	}
	return &Repos{repoRoot: repoRoot}, nil
}

// Diff returns a diff if differ between specified file's revision and latest on go tip.
func (hg *Repos) Diff(filepath, revision string) ([]byte, error) {
	cmd := exec.Command("hg", "diff", filepath, "-r", revision)
	cmd.Dir = hg.repoRoot

	return cmd.CombinedOutput()
}

type hglog struct {
	No  int `json:"n"`
	Rev string
}

type Status string

const (
	None     Status = "none"
	Ahead           = "ahead"
	Same            = "same"
	Outdated        = "outdated"
)

//
func (hg *Repos) Check(tag, filepath, revision string) (st Status, diff int, err error) {
	st = None
	diff = 0
	err = nil

	cmd := exec.Command("hg", "log", "-b", "default", "--template", `\{"n":{rev},"rev":"{node}"},`, "-r", tag+":"+revision, filepath)
	cmd.Dir = hg.repoRoot

	b, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	if len(b) == 0 {
		return
	}

	json_string := "["
	json_string = json_string + string(b[0:len(b)-1]) + "]"

	var logs []hglog
	json.Unmarshal([]byte(json_string), &logs)

	if len(logs) != 1 {
		for n, l := range logs {
			if strings.Contains(l.Rev, revision) {
				diff = n
				break
			}
		}
		if (logs[0].No - logs[1].No) < 0 {
			// ahead
			st = Ahead
		} else {
			// outdated
			st = Outdated
		}

	} else {
		st = Same
	}

	return
}

// UpdateRepo perform the hg pull;hg update tip.
func (hg *Repos) UpdateRepos() error {
	if err := hg.pull(); err != nil {
		return err
	}
	if err := hg.updateTip(); err != nil {
		return err
	}
	return nil
}

// Pull perform the hg pull.
func (hg *Repos) pull() error {
	cmd := exec.Command("hg", "pull")
	cmd.Dir = hg.repoRoot

	_, err := cmd.CombinedOutput()
	return err
}

// UpdateTip perform the hg update tip.
func (hg *Repos) updateTip() error {
	cmd := exec.Command("hg", "update", "tip")
	cmd.Dir = hg.repoRoot

	_, err := cmd.CombinedOutput()
	return err
}
