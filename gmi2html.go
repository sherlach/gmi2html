package main

import (
	"flag"
	"fmt"
	"html"
	"os"
)

func main() {
	var lexer = Lex(os.Stdin)

	var preformatted = false
	var listing = false

	//flag variables
	var title string

	//flag declaration
	flag.StringVar(&title, "t", "untitled", "Give your html document a title. Default is `untitled`")

	flag.Parse()
	initial(title)
	var lineNumber = 0
	for item := range lexer {
		lineNumber++
		/* I find this ugly, but there was not clean way
		to integrate this with the existing switch statement */
		if listing && item.Type != LineList {
			fmt.Println("</ul>")
			listing = false
		}
		switch item.Type {
		case LineError:
			fmt.Printf("Error %d: %s\n", lineNumber, item.Value)
		case LineText:
			if item.Value == "" {
				fmt.Println("<br />")
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
				fmt.Println("</pre>")
			} else {
				fmt.Printf("<pre>")
			}
			preformatted = !preformatted
		case LinePreformatted:
			fmt.Printf("%s\n", html.EscapeString(item.Value))
		}
	}
	end()
}

func initial(title string) {
	fmt.Println("<!DOCTYPE HTML>")
	fmt.Println("<html lang=\"en\">")
	fmt.Println("<meta charset=\"utf-8\">")
	fmt.Println("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">")

	fmt.Println("<!--This webpage has been auto generated by sherlach's gmi2html (https://github.com/sherlach/gmi2html) Please report any errors found!)-->")
	fmt.Println("<head>")
	fmt.Printf("<title>%s</title>\n", title)
	fmt.Println("</head>")
	style()
	fmt.Println("<body>")
}

func style() {
	fmt.Printf("<style>\n\t@media (prefers-color-scheme: dark){\n\t\tbody {color:#fff;background:#000}\n\t\ta:link {color:#9cf}\n\t\ta:hover, a:visited:hover {color:#cef}\n\t\ta:visited {color:#c9f}\n\t}\n\tbody {\n\t\tmargin:1em auto;\n\t\tmax-width:40em;\n\t\tpadding:0 .62em;\n\t\tfont:1.2em/1.62 sans-serif;\n\t}\n\th1,h2,h3 {\n\t\tline-height:1.2;\n\t}\n\t@media print {\n\t\tbody {\n\t\t\tmax-width:none\n\t\t}\n\t}\n</style>\n")
	/*
	   <style>\n

	   	\t@media (prefers-color-scheme: dark){\n
	   		\t\tbody {color:#fff;background:#000}\n
	   		\t\ta:link {color:#9cf}\n
	   		\t\ta:hover, a:visited:hover {color:#cef}\n
	   		\t\ta:visited {color:#c9f}\n
	   	\t}\n
	   	\tbody {\n
	   		\t\tmargin:1em auto;\n
	   		\t\tmax-width:40em;\n
	   		\t\tpadding:0 .62em;\n
	   		\t\tfont:1.2em/1.62 sans-serif;\n
	   	\t}\n
	   	\th1,h2,h3 {\n
	   		\t\tline-height:1.2;\n
	   	\t}\n
	   	\t@media print{\n
	   		\t\tbody {\n
	   			\t\t\tmax-width:none\n
	   		\t\t}\n
	   	\t}\n
	   </style>\n
	*/
}

func end() {
	fmt.Println("</body>")
	fmt.Println("</html>")
}
