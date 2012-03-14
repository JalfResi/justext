package gojustext

import(
	"io"
	"io/ioutil"
	"log"
	"os"
	"utf8"
	"html"
	"fmt"
	"strings"
)

type Reader struct {
	LengthLow          int
	LengthHigh         int
	StoplistLanguage   string
	StopwordsLow       float64
	StopwordsHigh      float64
	MaxLinkDensity     float64
	MaxHeadingDistance int
	NoHeadings         bool
	r                  io.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		LengthLow:          70,
		LengthHigh:         200,
		StoplistLanguage:   "English",
		StopwordsLow:       0.30,
		StopwordsHigh:      0.32,
		MaxLinkDensity:     0.2,
		MaxHeadingDistance: 200,
		NoHeadings:         false,	
		r:                  r,
	}
}

func (r *Reader) ReadAll() ([]*Paragraph, os.Error) {
	in, err := ioutil.ReadAll(r.r)
	if err != nil {
		return nil, err
	}

	stoplist, err := getStoplist(r.StoplistLanguage)
	if err != nil {
		return nil, err
	}

	root := preprocess(utf8.NewString(string(in)).String(), "utf-8", "utf-8", "errors")
	if root == nil {
		log.Fatal("Preprocess has resulted in nil")
	}

	htmlSource := nodesToString(root)
	if len(htmlSource) == 0 {
		log.Fatal("MAIN: perprocess has returned an empty string")
	}

	p, err := paragraphObjectModel(htmlSource)
	if err != nil {
		return nil, err
	}
	if p == nil {
		log.Println(htmlSource)
		log.Fatal("MAIN: P is nil", err)
	}

	classifyParagraphs(p, stoplist, r.LengthLow, r.LengthHigh, r.StopwordsLow, r.StopwordsHigh, r.MaxLinkDensity, r.NoHeadings)
	reviseParagraphClassification(p, r.MaxHeadingDistance)

	// NOTE: Might be best if this reader returned []paragraphs
	// We can then have a Writer that takes []paragraphs and outputs it as HTML.
	// The Writer could then have an option specifying detailed out, etc;
	// and the Reader would be responsible for classification, the Writer
	// for parsing the []paragraphs for output.
	//
	// The example app would then simply use copy to push the stdin to stdout
	// e.g. io.Copy(writer, reader)

	return p, nil	
}

func dumpNodes(n *html.Node, tab int, exploreChildNodes bool) string {
	var childNodes string = ""
	if exploreChildNodes == true {
		if len(n.Child)>0 {
			for _, c := range n.Child {
				childNodes = fmt.Sprintf("%s%s\n", childNodes, dumpNodes(c, tab+1, true))
			}
		}
	}

	var t string
	switch(n.Type) {
	case html.ErrorNode:
		t = "Error"
	case html.TextNode:
		t = "Text"
	case html.DocumentNode:
		t = "Document"
	case html.ElementNode:
		t = "Element"
	case html.CommentNode:
		t = "Comment"
	case html.DoctypeNode:
		t = "Doctype"
	}

	tabStr := strings.Repeat(" ", tab)
	return fmt.Sprintf("%sType: %s\n%sData: %s\n%s", tabStr, t, tabStr, n.Data, childNodes)
}