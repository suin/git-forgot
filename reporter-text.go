package main

import (
	"fmt"
	"io"
)

type TextReporter struct {
	w io.Writer
}

func NewTextReporter(w io.Writer) *TextReporter {
	return &TextReporter{
		w: w,
	}
}

func (r *TextReporter) Status(path string, status GitStatus) error {
	if status.NeedsToPushAndCommit() {
		fmt.Fprintln(r.w, path+" needs to push and commit.")
	} else if status.NeedsToPush {
		fmt.Fprintln(r.w, path+" needs to push.")
	} else if status.NeedsToCommit {
		fmt.Fprintln(r.w, path+" needs to commit.")
	}
	return nil
}
