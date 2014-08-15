package main

import (
	"os/exec"
)

type TerminalNotifierReporter struct {
}

func NewTerminalNotifierReporter() *TerminalNotifierReporter {
	return &TerminalNotifierReporter{}
}

func (r *TerminalNotifierReporter) Status(path string, status GitStatus) error {
	if status.NothingToDo() {
		return nil
	}

	var title string
	if status.NeedsToPushAndCommit() {
		title = "Need to git push and commit"
	} else if status.NeedsToPush {
		title = "Need to git push"
	} else if status.NeedsToCommit {
		title = "Need to git commit"
	}

	exec.Command("terminal-notifier", "-title", title, "-message", path).Run()
	return nil
}
