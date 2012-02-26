package gojustext

import(
	"io"
	"os"
)

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

func (w *Writer) outputDefault(paragraphs []*Paragraph) os.Error {
	b := []byte
	w.w.Write(b)

	return nil
}

func (w *Writer) outputDetailed(paragraphs []*Paragraph) os.Error {
	b := []byte
	w.w.Write(b)

	return nil
}