# Convert GMI files to HTML

[gmi2html][2] is a [text/gemini][0] to HTML converter written in Go:

    $ go get git.sr.ht/~henryprecheur/gmi2html
    $ go install git.sr.ht/~henryprecheur/gmi2html

Its design is inspired from Rob Pike’s talk: [Lexical Scanning in Go][1]. The
state of the lexer is kept in a callback, this neat trick simplifies the
lexer, and makes it more efficient. gmi2html reads its input from stdin and
writes the result to stdout, and there’s no flag:

    $ gmi2html < input.gmi > output.html

It doesn’t support any extension like list, and heading yet.

[0]: https://gemini.circumlunar.space/docs/specification.html
[1]: https://talks.golang.org/2011/lex.slide#1
[2]: https://git.sr.ht/~henryprecheur/gmi2html
