package gojustext

import (
	"fmt"
	"log"
	"os"
	"utf8"
)

const(
	LENGTH_LOW_DEFAULT int = 70
	LENGTH_HIGH_DEFAULT int = 200
	STOPWORDS_LOW_DEFAULT float64 = 0.30
	STOPWORDS_HIGH_DEFAULT float64 = 0.32
	MAX_LINK_DENSITY_DEFAULT float64 = 0.2
	MAX_HEADING_DISTANCE_DEFAULT = 200
	NO_HEADINGS_DEFAULT bool = false
)

// Encoding stuff...
// Do we need to do this?
// We probably should, just to be feature complete.
//
//func Justext(htmlText string, stoplist map[string]bool, lengthLow int, lengthHigh int, stopwordsLow float64, stopwordsHigh float64, maxLinkDensity float64, maxHeadingDistance int, noHeadings bool, encoding string, defaultEncoding string, encErrors string) []*Paragraph {


type JusText struct {
	HtmlText string
	Stoplist map[string]bool
	LengthLow int
	LengthHigh int
	StopwordsLow float64
	StopwordsHigh float64
	MaxLinkDensity float64
	MaxHeadingDistance int
	NoHeadings bool
}

func NewJusText() *JusText {
	return &JusText{
		LengthLow: LENGTH_LOW_DEFAULT,
		LengthHigh: LENGTH_HIGH_DEFAULT,
		StopwordsLow: STOPWORDS_LOW_DEFAULT,
		StopwordsHigh: STOPWORDS_HIGH_DEFAULT,
		MaxLinkDensity: MAX_LINK_DENSITY_DEFAULT,
		MaxHeadingDistance: MAX_HEADING_DISTANCE_DEFAULT,
		NoHeadings: NO_HEADINGS_DEFAULT,
	}
}

func (jusText *JusText)Process(htmlText string, stoplist map[string]bool) (string, os.Error) {
	
	root := preprocess(utf8.NewString(htmlText).String(), "utf-8", "utf-8", "errors")
	p, err := ParagraphObjectModel(nodesToString(root))
	if err != nil {
		log.Fatal(err)
	}
	if p == nil {
		log.Fatal("MAIN: P is nil", err)
	}

	classifyParagraphs(p, stoplist, jusText.LengthLow, jusText.LengthHigh, jusText.StopwordsLow, jusText.StopwordsHigh, jusText.MaxLinkDensity, jusText.NoHeadings)
	reviseParagraphClassification(p, jusText.MaxHeadingDistance)

	return OutputDefault(p, true), nil
}
/*
func Justext(htmlText string, stoplist map[string]bool, lengthLow int, lengthHigh int, stopwordsLow float64, stopwordsHigh float64, maxLinkDensity float64, maxHeadingDistance int, noHeadings bool) []*Paragraph {
	
	htmlText = utf8.NewString(htmlText).String()

	root := preprocess(htmlText, "utf-8", "utf-8", "errors")
	htmlText = nodesToString(root)

	p, err := ParagraphObjectModel(htmlText)
	if err != nil {
		log.Fatal(err)
	}
	if p == nil {
		log.Fatal("MAIN: P is nil", err)
	}

	classifyParagraphs(p, stoplist, lengthLow, lengthHigh, stopwordsLow, stopwordsHigh, maxLinkDensity, noHeadings)
	reviseParagraphClassification(p, maxHeadingDistance)

	return p
}
*/