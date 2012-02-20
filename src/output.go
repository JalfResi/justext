package gojustext

import(
	"fmt"
	"strings"
)

func OutputDefault(paragraphs []*Paragraph, noBoilerplate bool) {
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

func OutputDetailed(paragraphs []*Paragraph) (output string) {
	for _, paragraph := range paragraphs {
		output = fmt.Sprintf("%s<p class=\"%s\" cfclass=\"%s\" heading=\"%t\"> %s", output, paragraph.Class, paragraph.CfClass, paragraph.Heading, strings.TrimSpace(paragraph.Text))
	}

	return output
}

func OutputDebug(paragraphs []*Paragraph) (output string) {
	var counter int = 1
	for _, paragraph := range paragraphs {
		output = fmt.Sprintf("%s<p class=\"%s\" cfclass=\"%s\" heading=\"%t\"> %s", output, paragraph.Class, paragraph.CfClass, paragraph.Heading, strings.TrimSpace(paragraph.Text))

		var rows []string
		var tdtemp string = "<tr><td>%s</td><td>%s</td></tr>"
		var tdtempbool string = "<tr><td>%s</td><td>%t</td></tr>"
		var tdtempint string = "<tr><td>%s</td><td>%d</td></tr>"
		var tdtempfloat string = "<tr><td>%s</td><td>%f</td></tr>"
		rows = append(rows, fmt.Sprintf(tdtemp, "final class", paragraph.Class))
		rows = append(rows, fmt.Sprintf(tdtemp, "context-free class", paragraph.CfClass))
		rows = append(rows, fmt.Sprintf(tdtempbool, "heading", paragraph.Heading))
		rows = append(rows, fmt.Sprintf(tdtempint, "length (in characters)", len(paragraph.Text)))
		rows = append(rows, fmt.Sprintf(tdtempint, "number of characters with links", paragraph.LinkedCharCount))
		rows = append(rows, fmt.Sprintf(tdtempfloat, "link density", paragraph.LinkDensity))
		rows = append(rows, fmt.Sprintf(tdtempint, "number of words", paragraph.WordCount))
		rows = append(rows, fmt.Sprintf(tdtempint, "number of stop words", paragraph.StopwordCount))
		rows = append(rows, fmt.Sprintf(tdtempfloat, "stop word density", paragraph.StopwordDensity))
		rows = append(rows, fmt.Sprintf("<tr><td colspan=\"2\">%s</td></tr>", paragraph.DomPath))

		output = fmt.Sprintf("%s<div id=\"debug_%d\"><table>%s</table></div>", output, counter, strings.Join(rows, "\n"))
		counter++
	}

	output = fmt.Sprintf("<style>.good{background-color:green;} .bad{background-color:red;}</style><div style=\"width:500px\">%s</div>", output)

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
