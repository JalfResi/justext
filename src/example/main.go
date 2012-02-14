package main

import (
	//"fmt"
	"io/ioutil"
	"os"
	"ourscienceistight/gojustext"
)

func main() {

	in, _ := ioutil.ReadAll(os.Stdin)
	htmlStr := string(in)


	htmlStr = gojustext.Process(htmlStr)
	//fmt.Println(htmlStr)
}

