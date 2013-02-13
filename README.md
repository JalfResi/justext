gojustext
=========

A Go package that implements the JusText boilerplate removal algorithm (http://code.google.com/p/justext/)

## Install

    go get github.com/JalfResi/gojustext

And import:

    import "justext"

Supports all stoplist files available at http://code.google.com/p/justext/source/browse/#svn%2Ftrunk%2Fjustext%2Fstoplists

GoJustext expects valid HTML; it is your responsability to ensure that valid HTML is passed to GoJustext.
