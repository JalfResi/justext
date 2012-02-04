package gojustext

import (
	"log"
	"utf8"
)

func Process(htmlStr string) string {
	
	htmlStr = utf8.NewString(htmlStr).String()

	root := preprocess(htmlStr, "utf-8", "utf-8", "errors")
	htmlStr = nodesToString(root)

	p, err := ParagraphObjectModel(htmlStr)
	if err != nil {
		log.Fatal(err)
	}

	if p == nil {
		log.Fatal("P is nil")
	}

	htmlStr = nodesToString(root)
	//fmt.Println(htmlStr)

/*
	for _, para := range p {
		fmt.Printf("%s\n\tText: %s\n\tTagCount: %d\n\tWordCount: %d\n\tLinkedCharCount: %d\n", para.DomPath, para.Text, para.TagCount, para.WordCount, para.LinkedCharCount)
	}
*/
	return htmlStr
}


