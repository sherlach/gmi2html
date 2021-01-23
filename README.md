This program converts gemtext formatted documents (.gmi) to html.

It is a fork of Henry Precheur's program of the same name. 
(git.sr.ht/~henryprecheur/gmi2html)

I have added support for additional aspects of gemtext such as quoting,
lists, and headings. Style will be stolen from the best motherf$%@ing 
website. I'm not going to add a link - if you know, you know.

```
$ go get github.com/sherlach/gmi2html
$ go install github.com/sherlach/gmi2html
```

USAGE:

gmi2html -t TITLE < input.gmi > output.html

(TITLE is the title you want the HTML page to have. Given `untitled` as default.)

DISCLAIMERS:

1. PLEASE edit your html output after generated, there will be some errors, and
gemini links will still be in gemini form - you may want to edit them.
2. Before you use the program, make sure you read the files (especially gmi2html.go) and change anything you want to change (eg the initial and style functions). Submit a Report if you need help (but I doubt you will...)

TODO

- Refactoring desperately needed, os. module sometimes used, Printf sometimes used, for example. TrimPrefix instead of TrimLeft. AND, there's a lot of self-repetition
- Error handling
- Clean command line args
- HTML closing tag at the end
- favicon support??
- footers including "last updated"

to generate a full HTML document for you including doc and meta tags,
and specify a title (which isn't a thing in Gemini)
