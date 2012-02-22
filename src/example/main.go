package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"ourscienceistight/gojustext"
)

func main() {
	in, _ := ioutil.ReadAll(os.Stdin)

	j := gojustext.NewJusText()
	fmt.Println(j.Process(string(in)))

/*
	paragraphs := gojustext.Justext(string(in), stoplist, gojustext.LENGTH_LOW_DEFAULT, gojustext.LENGTH_HIGH_DEFAULT, gojustext.STOPWORDS_LOW_DEFAULT, gojustext.STOPWORDS_HIGH_DEFAULT, gojustext.MAX_LINK_DENSITY_DEFAULT, gojustext.MAX_HEADING_DISTANCE_DEFAULT, gojustext.NO_HEADINGS_DEFAULT)
	
	fmt.Println(gojustext.OutputDetailed(paragraphs))
*/
}

