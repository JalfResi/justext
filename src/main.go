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

/*
func Justext(htmlText string, stoplist []string, lengthLow int, lengthHigh int, stopwordsLow float64, stopwordsHigh float64, maxLinkDensity float64, maxHeadingDistance int, noHeadings bool, encoding string, defaultEncoding string, encErrors string) []*Paragraph {
	
}
*/

func Process(htmlStr string) string {
	
	htmlStr = utf8.NewString(htmlStr).String()

	root := preprocess(htmlStr, "utf-8", "utf-8", "errors")
	htmlStr = nodesToString(root)

	p, err := ParagraphObjectModel(htmlStr)
	if err != nil {
		log.Fatal(err)
	}
	if p == nil {
		log.Fatal("MAIN: P is nil", err)
	}

	classifyParagraphs(p, stoplists["English"], LENGTH_LOW_DEFAULT, LENGTH_HIGH_DEFAULT, STOPWORDS_LOW_DEFAULT, STOPWORDS_HIGH_DEFAULT, MAX_LINK_DENSITY_DEFAULT, NO_HEADINGS_DEFAULT)

	reviseParagraphClassification(p, MAX_HEADING_DISTANCE_DEFAULT)

	//outputDefault(p, true)
	out := outputDebug(p)
	fmt.Println(out)

	//dump(p)

	return htmlStr
}

func dump(p []*Paragraph) {
	for _, para := range p {
		fmt.Printf("%s\n\tText: %s\n\tTagCount: %d\n\tWordCount: %d\n\tLinkedCharCount: %d\n\tStopword Count: %d\n\tStopword Density: %f\n\tLink Density: %f\n\tHeading: %t\n\tClass: %s\n\tCfClass: %s\n", para.DomPath, para.Text, para.TagCount, para.WordCount, para.LinkedCharCount, para.StopwordCount, para.StopwordDensity, para.LinkDensity, para.Heading, para.Class, para.CfClass)
	}	
}

