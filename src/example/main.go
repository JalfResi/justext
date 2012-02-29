package main

import (
	"flag"
	"log"
	"os"
	"ourscienceistight/gojustext"
)

var(
	outputMode *bool = flag.Bool("detailed", false, "Generates HTML with detailed results")
)

func main() {
	flag.Parse()

	jusText := gojustext.NewReader(os.Stdin)
	jusText.NoHeadings = true
	p, err := jusText.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	j := gojustext.NewWriter(os.Stdout)

	if *outputMode {
		j.Mode = gojustext.MODE_DETAILED
	}

	j.WriteAll(p)
}

