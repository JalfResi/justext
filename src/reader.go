package gojustext

import(
	"io"
	"io/ioutil"
	"log"
	"os"
	"utf8"
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

func (r *Reader) ReadAll() (htmlStr string, err os.Error) {
	in, err := ioutil.ReadAll(r.r)
	if err != nil {
		return "", err
	}

	stoplist, err := getStoplist(r.StoplistLanguage)
	if err != nil {
		return "", err
	}

	root := preprocess(utf8.NewString(string(in)).String(), "utf-8", "utf-8", "errors")
	p, err := paragraphObjectModel(nodesToString(root))
	if err != nil {
		return "", err
	}
	if p == nil {
		log.Fatal("MAIN: P is nil", err)
	}

	classifyParagraphs(p, stoplist, r.LengthLow, r.LengthHigh, r.StopwordsLow, r.StopwordsHigh, r.MaxLinkDensity, r.NoHeadings)
	reviseParagraphClassification(p, r.MaxHeadingDistance)

	// NOTE: Might be best if this reader returned []paragraphs
	// We can then have a Writer that takes []paragraphs and outputs it as HTML.
	// The Writer could then have an option specifying detailed out, etc;
	// and the Reader would be responsible for classification, the Writer
	// for parsing the []paragraphs for output.

	return OutputDefault(p, true), nil	
}