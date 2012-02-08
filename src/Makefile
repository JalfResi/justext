include $(GOROOT)/src/Make.inc

TARG=ourscienceistight/gojustext
GOFILES=\
	main.go\
	preprocess.go\
	paragraphObjectModel.go\
	htmlRenderer.go\
	classifyParagraphs.go\
	stoplists/English.go\

include $(GOROOT)/src/Make.pkg

stoplists/English.go:
	php go-mapdata.php -fhttp://justext.googlecode.com/svn/trunk/justext/stoplists/English.txt -o"stoplists/English.go" -n"EnglishStoplist" -p"gojustext"