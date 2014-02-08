package gotip

import (
	"log"
	"os"
	"os/exec"
)

type Repo struct {
	repoRoot string
}

// AttachRepos is handle of hg command.
// repoRoot is represent hg repository root.
func AttachRepos(repoRoot string) (*Repo, error) {
	if _, err := os.Stat(repoRoot); err != nil {
		return nil, err
	}
	if _, err := os.Stat(repoRoot + "/.hg"); err != nil {
		return nil, err
	}
	return &Repo{repoRoot: repoRoot}, nil
}

// Diff returns a diff if differ between specified file's revision and latest on go tip.
func (hg *Repo) Diff(filepath, revision string) ([]byte, error) {
	cmd := exec.Command("hg", "diff", filepath, "-r", revision)
	cmd.Dir = hg.repoRoot

	return cmd.CombinedOutput()
}

// UpdateRepo perform the hg pull;hg update tip.
func (hg *Repo) UpdateRepo() error {
	if err := hg.pull(); err != nil {
		return err
	}
	if err := hg.updateTip(); err != nil {
		return err
	}
	return nil
}

// Pull perform the hg pull.
func (hg *Repo) pull() error {
	cmd := exec.Command("hg", "pull")
	cmd.Dir = hg.repoRoot

	_, err := cmd.CombinedOutput()
	return err
}

// UpdateTip perform the hg update tip.
func (hg *Repo) updateTip() error {
	cmd := exec.Command("hg", "update", "tip")
	cmd.Dir = hg.repoRoot

	_, err := cmd.CombinedOutput()
	return err
}

// IsLatest returns true if file's revision is latest on go tip
func (hg *Repo) IsLatest(filepath, revision string) (bool, error) {
	b, err := hg.Diff(filepath, revision)
	if err != nil {
		return false, err
	}

	if len(b) != 0 {
		return false, nil
	}

	return true, nil
}

func run(c *exec.Cmd) {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
