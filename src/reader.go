package gojustext

import(
	"io"
	"io/ioutil"
	"os"
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
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	stoplist, err := GetStoplist(jusText.StoplistLanguage)
	if err != nil {
		return nil, err
	}

	root := preprocess(utf8.NewString(htmlText).String(), "utf-8", "utf-8", "errors")
	p, err := ParagraphObjectModel(nodesToString(root))
	if err != nil {
		return nil, err
	}
	if p == nil {
		log.Fatal("MAIN: P is nil", err)
	}

	classifyParagraphs(p, stoplist, jusText.LengthLow, jusText.LengthHigh, jusText.StopwordsLow, jusText.StopwordsHigh, jusText.MaxLinkDensity, jusText.NoHeadings)
	reviseParagraphClassification(p, jusText.MaxHeadingDistance)

	return OutputDefault(p, true), nil	
}