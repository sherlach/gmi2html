This program converts gemtext formatted documents (.gmi) to html.

It is a fork of Henry Precheur's program of the same name. 
(git.sr.ht/~henryprecheur/gmi2html)

I have added support for additional aspects of gemtext such as quoting,
lists, and headings. Style will be stolen from the best motherf$%@ing 
website. I'm not going to add a link - if you know, you know.

(Work in Progress)


TODO

- Refactoring desperately needed, os. module sometimes used, Printf sometimes used, for example. TrimPrefix instead of TrimLeft. AND, there's a lot of self-repetition
- Heading "Doctype html" et cetera
- Similarly, Style from the BMW
- Quotes
- Headers have strange behaviour, '### ' and '## ' are expected, but the # prefixis triggering without whitespace

- Command line args:

to generate a full HTML document for you including doc and meta tags,
and specify a title (which isn't a thing in Gemini)
