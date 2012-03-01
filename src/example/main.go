package main

import (
	"flag"
	"log"
	"os"
	"ourscienceistight/gojustext"
)

var(
	noHeadings *bool = flag.Bool("no-headings", false, "Ignores headings in output")
	outputMode *bool = flag.Bool("detailed", false, "Generates HTML with detailed results")
	debugMode  *bool = flag.Bool("debug", false, "Log debug messages")
)

func main() {
	flag.Parse()

	jusText := gojustext.NewReader(os.Stdin)
	jusText.NoHeadings = *noHeadings
	p, err := jusText.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	j := gojustext.NewWriter(os.Stdout)
	if *outputMode {
		j.Mode = gojustext.MODE_DETAILED
	}

	if *debugMode {
		j.OutputDebug(p)
	}

	j.WriteAll(p)
}

