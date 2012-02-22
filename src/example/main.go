package main

import (
	"fmt"
	"log"
	"os"
	"ourscienceistight/gojustext"
)

func main() {

	jusText := gojustext.NewReader(os.Stdin)
	jusText.NoHeadings = true
	output, err := jusText.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

