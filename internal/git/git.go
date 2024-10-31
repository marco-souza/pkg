package git

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Clone(url, name string) error {
	url = parseURL(url)
	gitCommand := "clone " + url

	if strings.Trim(name, " ") == "" {
		repoSplit := strings.Split(url, "/")
		name = repoSplit[len(repoSplit)-1]
	}

	gitCommand += " " + name

	_, err := exec.LookPath("git")
	if errors.Is(err, exec.ErrDot) {
		err = errors.New("git command not found")
	}
	if err != nil {
		return err
	}

	// if folder exists, remove it
	if err := removePath(name); err != nil {
		return errors.Join(errors.New("error removing folder"), err)
	}

	c := exec.Command("git", strings.Split(gitCommand, " ")...)
	if err := c.Run(); err != nil {
		return errors.Join(errors.New("error running git command: git "+gitCommand), err)
	}

	gitPath := path.Join(name, ".git")
	if err := removePath(gitPath); err != nil {
		return errors.Join(errors.New("error removing .git"), err)
	}

	return nil
}

func removePath(path string) error {
	if _, err := os.Stat(path); err == nil {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}

func parseURL(repo string) string {
	if !strings.Contains(repo, "github.com") {
		// TODO: support different providers
		repo = "github.com/" + repo
	}

	if !strings.HasPrefix(repo, "https://") {
		// TODO: support different protocols
		repo = "https://" + repo
	}

	return repo
}
