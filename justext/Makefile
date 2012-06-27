include $(GOROOT)/src/Make.inc

TARG=ourscienceistight/gojustext
GOFILES=\
	preprocess.go\
	paragraphObjectModel.go\
	htmlRenderer.go\
	classifyParagraphs.go\
	reviseParagraphClassification.go\
	reader.go\
	writer.go\
	stoplists.go\
	stoplists/*.go\
	defaultTemplate.go\
	detailedTemplate.go\

include $(GOROOT)/src/Make.pkg

stoplists/%.go:
	wget -Ostoplists/$*.txt http://justext.googlecode.com/svn/trunk/justext/stoplists/$*.txt
	go-bindata -i=stoplists/$*.txt -o="stoplists/$*.go" -f="$*Stoplist" -p="gojustext"

defaultTemplate.go:
	go-bindata -i=default.template -o="defaultTemplate.go" -f="DefaultTemplate" -p="gojustext"

detailedTemplate.go:
	go-bindata -i=detailed.template -o="detailedTemplate.go" -f="DetailedTemplate" -p="gojustext"
