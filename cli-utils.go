package main

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/koyachi/go-term-ansicolor/ansicolor"
	"go/doc"
	"strings"
)

func printError(message string) {
	if log.IsTerminal() {
		message = wrapString(message, 70)
		message = ansicolor.Red(message)
	}
	fmt.Println(message)
}

func wrapString(text string, length int) string {
	if strings.ContainsAny(text, "\r\n") {
		return text
	}
	buf := bytes.NewBufferString("")
	doc.ToText(buf, text, "", "", length)
	return buf.String()
}
