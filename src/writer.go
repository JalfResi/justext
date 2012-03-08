package gojustext

import(
	"log"
	"io"
	"os"
	"template"
	"strings"
)

// NOTE:
// Make a new type:
//  type JusText []paragraphs


const(
	MODE_DEFAULT  = 1
	MODE_DETAILED = 2
)

type Writer struct {
	Mode          int
	NoBoilerplate bool
	w             io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Mode:          MODE_DEFAULT,
		NoBoilerplate: true,
		w:             w,
	}
}

func (w *Writer) WriteAll(paragraphs []*Paragraph) os.Error {
	switch(w.Mode) {
		case MODE_DEFAULT:
			return w.outputDefault(paragraphs)
			break

		case MODE_DETAILED:
			return w.outputDetailed(paragraphs)
			break

		default:
			return os.NewError("Unrecognised mode")
	}

	return nil
}

func IsGood(args ...interface{}) (result bool) {
	result = true
	for _, val := range args {
		if val != "good" {
			result = false
			return
		}
	}
	return
}

func (w *Writer) outputDefault(paragraphs []*Paragraph) os.Error {
	templateData, err := DefaultTemplate()
	if err != nil {
		log.Fatal(err)
	}

	t := template.New("default")
	t.Funcs(template.FuncMap{"TrimSpace": strings.TrimSpace})
	t.Funcs(template.FuncMap{"IsGood": IsGood})

	templ, err := t.Parse(string(templateData))
	if err != nil {
		log.Fatal(err)
	}

	var data = struct {
		Paragraphs []*Paragraph
		NoBoilerplate bool
	}{paragraphs, w.NoBoilerplate}

	return templ.Execute(w.w, data)
}

func (w *Writer) outputDetailed(paragraphs []*Paragraph) os.Error {
	templateData, err := DetailedTemplate()
	if err != nil {
		log.Fatal(err)
	}
	
	t := template.New("detailed")
	t.Funcs(template.FuncMap{"TrimSpace": strings.TrimSpace})
	
	templ, err := t.Parse(string(templateData))
	//templ, err := t.ParseFile("/Users/bendavies/Projects/gojustext/src/detailed.template")
	if err != nil {
		log.Fatal(err)
	}

	var data = struct {
		Paragraphs []*Paragraph
	}{paragraphs}

	return templ.Execute(w.w, data)
}

func (w *Writer) OutputDebug(paragraphs []*Paragraph) {
	for _, paragraph := range paragraphs {
		log.Println(paragraph.DomPath)
		log.Println("\tfinal class: ", paragraph.Class)
		log.Println("\tcontext-free class: ", paragraph.CfClass)
		log.Println("\theading: ", paragraph.Heading)
		log.Println("\tlength (in characters): ", len(paragraph.Text))
		log.Println("\tnumber of characters with links: ", paragraph.LinkedCharCount)
		log.Println("\tlink density: ", paragraph.LinkDensity)
		log.Println("\tnumber of words: ", paragraph.WordCount)
		log.Println("\tnumber of stop words: ", paragraph.StopwordCount)
		log.Println("\tstop word density: ", paragraph.StopwordDensity)
	}
}

// TO-DO:
// Need an output feature that returns a de-duped space separated text file of all the 
// words in the output document sans-boilerplate. Also needs option to exclude stoplist
// words from that output too.

// TO-DO:
// Need an output feature that returns the content of a stop list (or do we just make 
// the function getStoplist public? Might be a lot easier...)

/*
func (w *Writer) outputKrdwrd(paragraphs []*Paragraph) (output string) {
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
*/