package gojustext

func classifyParagraphs(paragraphs []*Paragraph, stoplist map[string]bool, lengthLow int, lengthHigh int stopwordsLow float, stopwordsHigh float, maxLinkDensity float, noHeadings bool) {
	for _, paragraph := range paragraphs {
		var length int = len(paragraph.Text)
		var stopwordCount int = 0
		for _, word := range strings.Split(paragraph.Text, " ") {
			if _, ok := stoplist[word]; ok {
				stopwordCount += 1
			}
		}

		var stopwordDensity float = 0.0
		var linkDensity float = 0.0
		var wordCount int = paragraph.WordCount

		if wordCount > 0 {
			stopwordDensity = 1.0 * stopwordCount/wordCount
			linkDensity = float(paragraph.LinkedCharCount)/length
		}

		paragraph.StopwordCount = stopwordCount
		paragraph.StopwordDensity = stopwordDensity
		paragraph.LinkDensity = linkDensity

		var findHeadings *regexp.Regexp = regexp.MustCompile("(^h\d|\.h\d)")
		paragraph.Heading = bool(!noHeadings && findHeadings.FindString(paragraph.DomPath))

		var copyrightChar *regexp.Regexp = regexp.MustCompile("(\xa9|&copy)")
		var findSelect *regexp.Regexp = regexp.MustCompile("(^select|\.select)")

		if linkDensity > maxLinkDensity {
			paragraph.CfClass = "bad"
		} else if copyrightChar.FindString(paragraph.Text) {
			paragraph.CfClass = "bad"			
		} else if findSelect.FindString(paragraph.DomPath) {
			paragraph.CfClass = "bad"			
		} else {
			if length < lengthLow {
				if paragraph.LinkedCharCount > 0 {
					paragraph.CfClass = "bad"						
				} else {
					paragraph.CfClass = "short"			
				}
			} else {
				if stopwordDensity >= stopwordsHigh {
					if length > lengthHigh {
						paragraph.CfClass = "good"		
					} else {
						paragraph.CfClass = "neargood"			
					}
				} else {
					if stopwordDensity >= stopwordsLow {
						paragraph.CfClass = "neargood"
					} else {
						paragraph.CfClass = "bad"
					}
				}
			}
		}
	}
}