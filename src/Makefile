include $(GOROOT)/src/Make.inc

TARG=ourscienceistight/gojustext
GOFILES=\
	preprocess.go\
	paragraphObjectModel.go\
	htmlRenderer.go\
	classifyParagraphs.go\
	reviseParagraphClassification.go\
	output.go\
	reader.go\
	stoplists.go\
	stoplists/*.go\

include $(GOROOT)/src/Make.pkg

stoplists/%.go:
	wget -Ostoplists/$*.txt http://justext.googlecode.com/svn/trunk/justext/stoplists/$*.txt
	go-bindata -i=stoplists/$*.txt -o="stoplists/$*.go" -f="$*Stoplist" -p="gojustext"
