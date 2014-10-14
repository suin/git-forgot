package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
)

func checkAll(dirs []string, reporterNames []string, termAppName string) error {
	for _, dir := range dirs {
		err := check(dir, reporterNames, termAppName)
		if err != nil {
			return err
		}
	}

	return nil
}

func check(dir string, reporterNames []string, termAppName string) error {
	log.Debug("Check dir: " + dir)
	paths, err := filepath.Glob(dir)
	if err != nil {
		return err
	}

	var reporters = map[string]Reporter{
		"text":              NewTextReporter(os.Stdout),
		"terminal-notifier": NewTerminalNotifierReporter(termAppName),
	}

	reporter := NewReporterReporter()
	for _, name := range reporterNames {
		if r, ok := reporters[name]; ok {
			reporter.Append(r)
		}
	}

	for _, path := range paths {
		git, err := NewGit(path)
		if err == nil {
			status, err := git.Status()
			if err != nil {
				return err
			}
			reporter.Status(git.Path, status)
		}
	}

	return nil
}
