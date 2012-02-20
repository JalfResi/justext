package gojustext

import (
	"fmt"
	"log"
	"utf8"
)

const(
	MAX_LINK_DENSITY_DEFAULT float64 = 0.2
	LENGTH_LOW_DEFAULT int = 70
	LENGTH_HIGH_DEFAULT int = 200
	STOPWORDS_LOW_DEFAULT float64 = 0.30
	STOPWORDS_HIGH_DEFAULT float64 = 0.32
	NO_HEADINGS_DEFAULT bool = false
	MAX_HEADING_DISTANCE_DEFAULT = 200
)

// Encoding stuff...
// Do we need to do this?
// We probably should, just to be feature complete.
//
//func Justext(htmlText string, stoplist map[string]bool, lengthLow int, lengthHigh int, stopwordsLow float64, stopwordsHigh float64, maxLinkDensity float64, maxHeadingDistance int, noHeadings bool, encoding string, defaultEncoding string, encErrors string) []*Paragraph {


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

func dump(p []*Paragraph) {
	for _, para := range p {
		fmt.Printf("%s\n\tText: %s\n\tTagCount: %d\n\tWordCount: %d\n\tLinkedCharCount: %d\n\tStopword Count: %d\n\tStopword Density: %f\n\tLink Density: %f\n\tHeading: %t\n\tClass: %s\n\tCfClass: %s\n", para.DomPath, para.Text, para.TagCount, para.WordCount, para.LinkedCharCount, para.StopwordCount, para.StopwordDensity, para.LinkDensity, para.Heading, para.Class, para.CfClass)
	}	
}

