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
	var lineNumber = 0
	for item := range lexer {
		lineNumber += 1
		switch item.Type {
		case LineError:
			fmt.Printf("Error %d: %s\n", lineNumber, item.Value)
		case LineText:
			if item.Value == "" {
				io.WriteString(os.Stdout, "<br />\n")
			} else {
				fmt.Printf("<p>%s</p>\n", html.EscapeString(item.Value))
			}
		case LineLink:
			var text = item.Value
			if text == "" {
				text = item.URL.String()
			}
			text = html.EscapeString(text)
			var u = item.URL.String()
			fmt.Printf("<a href='%s'>%s</a><br>\n", u, text)
		case LineList:
			fmt.Printf("<li>%s</li>\n", html.EscapeString(item.Value))
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
