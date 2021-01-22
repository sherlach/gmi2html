package main

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var exampleURL, _ = url.Parse("gemini://example.com/")

func TestTextLine(t *testing.T) {
	var cases = []struct {
		Line   string
		Result Line
	}{
		{"", Line{LineText, "", nil}},
		{"abc", Line{LineText, "abc", nil}},
		{"=>", Line{LineError, "No URL after link prefix '=>'", nil}},
		{"=> \t", Line{LineError, "No URL after link prefix '=>'", nil}},
		{"=>gemini://example.com/", Line{LineLink, "", exampleURL}},
		{"=> gemini://example.com/", Line{LineLink, "", exampleURL}},
		{"=>\tgemini://example.com/", Line{LineLink, "", exampleURL}},
		{
			"=> gemini://example.com/ Gemini Link to Example.com",
			Line{LineLink, "Gemini Link to Example.com", exampleURL},
		},
		{
			"=> \t \t gemini://example.com/ \t \t Gemini\t \tLink to Example.com",
			Line{LineLink, "Gemini\t \tLink to Example.com", exampleURL},
		},
		{"```", Line{LinePreformattedToggle, "", nil}},
		{"```Alt Text", Line{LinePreformattedToggle, "Alt Text", nil}},
		{"=>", Line{LineError, "No URL after link prefix '=>'", nil}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v", c.Line), func(t *testing.T) {
			var _, result = TextLine(c.Line)
			assert.Equal(t, c.Result, result)
		})
	}
}

func TestPreformattedLine(t *testing.T) {
	var cases = []struct {
		Line   string
		Result Line
	}{
		{"", Line{LinePreformatted, "", nil}},
		{"abc", Line{LinePreformatted, "abc", nil}},
		{" \t", Line{LinePreformatted, " \t", nil}},
		{"```", Line{LinePreformattedToggle, "", nil}},
		{"```Alt Text", Line{LinePreformattedToggle, "Alt Text", nil}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v", c.Line), func(t *testing.T) {
			var _, result = PreFormattedLine(c.Line)
			assert.Equal(t, result, c.Result)
		})
	}
}

func TestLex(t *testing.T) {
	var cases = []struct {
		Text    string
		Results []Line
	}{
		{"", []Line{}},
		{"\n", []Line{{LineText, "", nil}}},
		{"foo\nbar", []Line{{LineText, "foo", nil}, {LineText, "bar", nil}}},
		{
			"foo\n=> gemini://example.com/",
			[]Line{
				{LineText, "foo", nil},
				{LineLink, "", exampleURL},
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.Text), func(t *testing.T) {
			var channel = Lex(strings.NewReader(c.Text))
			for number, line := range c.Results {
				var result Line
				select {
				case result = <-channel:
					assert.Equal(t, line, result, "At line %d", number)
				case <-time.After(time.Second):
					t.Fatalf("Not enough data after line %d", number)
				}
			}

		})
	}
}
