package main

import (
	"fmt"
	"html"
	"io"
	"os"
)

func main() {
	var lexer = Lex(os.Stdin)

	var preformatted = false
	var listing = false

	var lineNumber = 0
	for item := range lexer {
		lineNumber += 1
		/* I find this ugly, but there was not clean way
		to integrate this with the existing switch statement */
		if listing && item.Type != LineList {
			fmt.Println("</ul>\n")
			listing = false
		}
		switch item.Type {
		case LineError:
			fmt.Printf("Error %d: %s\n", lineNumber, item.Value)
		case LineText:
			if item.Value == "" {
				io.WriteString(os.Stdout, "<br />\n")
			} else {
				fmt.Printf("<p>%s</p>\n", html.EscapeString(item.Value))
			}
		case LineHeader:
			fmt.Printf("<h%s>%s</h%s>\n", item.HeadSize, item.Value, item.HeadSize)
		case LineLink:
			var text = item.Value
			if text == "" {
				text = item.URL.String()
			}
			text = html.EscapeString(text)
			var u = item.URL.String()
			fmt.Printf("<a href='%s'>%s</a><br>\n", u, text)
		case LineList:
			if !listing {
				fmt.Printf("<ul>\n")
			}
			fmt.Printf("<li>%s</li>\n", html.EscapeString(item.Value))
			listing = true
		case LineQuote:
			fmt.Printf("<blockquote>%s</blockquote>\n", item.Value)
		case LinePreformattedToggle:
			if preformatted {
				io.WriteString(os.Stdout, "</pre>\n")
			} else {
				io.WriteString(os.Stdout, "<pre>")
			}
			preformatted = !preformatted
		case LinePreformatted:
			io.WriteString(os.Stdout, html.EscapeString(item.Value))
			io.WriteString(os.Stdout, "\n")
		}
	}
}
