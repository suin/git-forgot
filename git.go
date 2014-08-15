package main

import (
	"bytes"
	"errors"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Git struct {
	Path string
}

func NewGit(path string) (*Git, error) {
	git := &Git{}

	if git.isGitProject(path) == false {
		return git, NotGitProjectError{path}
	}

	git.Path = path

	return git, nil
}

func (git *Git) Status() (GitStatus, error) {
	status := GitStatus{}
	workingDir, err := os.Getwd()
	if err != nil {
		return status, err
	}
	defer os.Chdir(workingDir)

	err = os.Chdir(git.Path)
	if err != nil {
		return status, err
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("git", "status", "-sb")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	outString := string(bytes.Trim(stdout.Bytes(), "\r\n"))
	log.Debug(stdout.String())
	if err != nil {
		log.Debug(stderr.String())
		return status, err
	}
	lines := strings.Split(outString, "\n")
	if len(lines) < 1 {
		return status, errors.New("Command output should has one or more lines")
	}

	needsToPush, err := regexp.MatchString(`\[ahead \d+\]`, lines[0])
	if err != nil {
		return status, err
	}

	status.NeedsToPush = needsToPush
	status.NeedsToCommit = len(lines) > 1

	return status, nil
}

func (git *Git) isGitProject(path string) bool {
	path += "/.git"
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() == false {
		return false
	}

	return true
}

type GitStatus struct {
	NeedsToPush   bool
	NeedsToCommit bool
}

func (status GitStatus) NothingToDo() bool {
	return status.NeedsToPush == false && status.NeedsToCommit == false
}

func (status GitStatus) NeedsToPushAndCommit() bool {
	return status.NeedsToPush && status.NeedsToCommit
}

type NotGitProjectError struct {
	Path string
}

func (e NotGitProjectError) Error() string {
	return "This is not git project: " + e.Path
}
