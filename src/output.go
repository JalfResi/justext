package gojustext

import(
	"fmt"
	"strings"
)

func outputDefault(paragraphs []*Paragraph, noBoilerplate bool) {
	for _, paragraph := range paragraphs {
		var tag string
		if paragraph.Class == "good" {
			if paragraph.Heading {
				tag = "h"
			} else {
				tag = "p"
			}
		} else {
			if noBoilerplate {
				continue
			} else {
				tag = "b"
			}
		}
		fmt.Printf("<%s> %s", tag, strings.TrimSpace(paragraph.Text))
	}
}

func outputDetailed(paragraphs []*Paragraph) (output string) {
	for _, paragraph := range paragraphs {
		output = fmt.Sprintf("%s<p class=\"%s\" cfclass=\"%s\" heading=\"%t\"> %s", output, paragraph.Class, paragraph.CfClass, paragraph.Heading, strings.TrimSpace(paragraph.Text))
	}

	return output
}

func outputKrdwrd(paragraphs []*Paragraph) (output string) {
	for _, paragraph := range paragraphs {
		var cls int
		if paragraph.Class == "good" || paragraph.Class == "neargood" {
			if paragraph.Heading {
				cls = 2
			} else {
				cls = 3
			}
		} else {
			cls = 1
		}
		for _, textNode := range paragraph.TextNodes {
			output = fmt.Sprintf("%s%i\t%s", output, cls, strings.TrimSpace(textNode))		
		}
	}

	return output
}
