package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type LineType int

const (
	LineError LineType = iota

	LineText
	LineLink
	LineList
	LinePreformattedToggle
	LinePreformatted
)

type Line struct {
	Type  LineType
	Value string
	URL   *url.URL // nil if not LineLink
}

type StateFunc func(line string) (StateFunc, Line)

const (
	preformattedToggle = "```"
	linkPrefix         = "=>"
	bulletPoint	   = "* "
)

/* This code has been taken from the original project, but I think the
strings.TrimPrefix method could be more suitable than the TrimLeft used.
I will leave this decision to a future refactor. */
/* XXX In addition, gemtext is normally written without line breaks. Parsing
gemtext blocks into lines may also need to be handled. */
func TextLine(line string) (StateFunc, Line) {
	if strings.HasPrefix(line, linkPrefix) {
		var (
			trimmed  = strings.TrimLeft(line[len(linkPrefix):], " \t")
			endofurl = strings.IndexAny(trimmed, " \t")
			rawurl   string
			alttext  string
		)
		if endofurl == -1 {
			rawurl = trimmed
		} else {
			rawurl = trimmed[:endofurl]
			alttext = strings.TrimLeft(trimmed[endofurl:], " \t")
		}
		if rawurl == "" {
			return TextLine, Line{Type: LineError, Value: "No URL after link prefix '=>'"}
		}
		var parsedurl, err = url.Parse(rawurl)

		if err != nil {
			return TextLine, Line{LineError, fmt.Sprintf("Invalid URL: '%s'", rawurl), nil}
		}
		return TextLine, Line{LineLink, alttext, parsedurl}
	}
	if strings.HasPrefix(line, bulletPoint) {
		trimmed := strings.TrimLeft(line[len(bulletPoint):], "\t")
		return TextLine, Line{Type: LineList, Value: trimmed}
	}
	if strings.HasPrefix(line, preformattedToggle) {
		return PreFormattedLine, Line{
			Type:  LinePreformattedToggle,
			Value: line[len(preformattedToggle):],
		}
	}

	return TextLine, Line{Type: LineText, Value: line}
}

func PreFormattedLine(line string) (StateFunc, Line) {
	if strings.HasPrefix(line, preformattedToggle) {
		return TextLine, Line{
			Type:  LinePreformattedToggle,
			Value: line[len(preformattedToggle):],
		}
	}
	return PreFormattedLine, Line{Type: LinePreformatted, Value: line}
}

func Lex(r io.Reader) chan Line {
	var output = make(chan Line)
	var scanner = bufio.NewScanner(r)
	var stateFunc = TextLine

	go func() {
		for scanner.Scan() {
			var result Line
			stateFunc, result = stateFunc(scanner.Text())
			output <- result
		}
		close(output)
	}()

	return output
}
