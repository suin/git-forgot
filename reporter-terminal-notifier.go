package main

import (
	"os/exec"
)

type TerminalNotifierReporter struct {
	termAppName string
}

func NewTerminalNotifierReporter(termAppName string) *TerminalNotifierReporter {
	return &TerminalNotifierReporter{termAppName}
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

	switch r.termAppName {
	case "iTerm":
		execute := " /usr/bin/osascript -e 'activate application \"iTerm\"' && " +
			"/usr/bin/osascript -e 'tell application \"System Events\" to keystroke \"t\" using command down' && " +
			"/usr/bin/osascript -e 'tell application \"iTerm\" to tell session -1 of current terminal to write text \"cd " +
			path + "\"' "
		exec.Command("terminal-notifier", "-title", title, "-execute", execute, "-message", path).Run()
	case "Terminal":
		execute := " /usr/bin/osascript -e 'activate application \"Terminal\"' && " +
			"/usr/bin/osascript -e 'tell application \"System Events\" to keystroke \"t\" using command down' && " +
			"osascript -e 'tell application \"Terminal\" to do script \"cd " + path + "\" in window 1 '"
		exec.Command("terminal-notifier", "-title", title, "-execute", execute, "-message", path).Run()
	default:
		exec.Command("terminal-notifier", "-title", title, "-message", path).Run()
	}

	return nil
}
