package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

var version string

func init() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "git-forgot"
	app.Usage = "Oops! I fogot pushing my commits!"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Debug mode.",
		},
		cli.StringFlag{
			Name:  "reporter, r",
			Value: "text",
			Usage: "You can use two or many reporters by seperating comma like `-r text,terminal-notifier`.",
		},
		cli.StringFlag{
			Name:  "terminal, t",
			Value: "text",
			Usage: "If you want to return to the terminal when click the notification  `-t iTerm` or `-t Terminal`.",
		},
	}
	app.Before = func(c *cli.Context) (err error) {
		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return
	}
	app.Action = func(c *cli.Context) {
		reportersString := c.String("reporter")
		reporterNames := strings.Split(reportersString, ",")
		termAppName := c.String("terminal")
		directoriesToCheck := []string{}

		if len(c.Args()) > 0 {
			directoriesToCheck = c.Args()
		} else {
			dir := os.Getenv("GIT_FORGOT_DIR")
			dirs := strings.Split(dir, " ")
			if len(dir) == 0 {
				printError("You have to specify directories to check, or set environment variables `GIT_FORGOT_DIR`")
				os.Exit(1)
			}
			directoriesToCheck = dirs
		}

		err := checkAll(directoriesToCheck, reporterNames, termAppName)
		if err != nil {
			printError(err.Error())
			os.Exit(1)
		}
	}
	app.Run(os.Args)
}
