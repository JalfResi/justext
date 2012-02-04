package gojustext

import (
	"fmt"
	"html"
	"log"
	"utf8"
)

func Process(htmlStr string) string {
	
	htmlStr = utf8.NewString(htmlStr).String()

	root := preprocess(htmlStr, "utf-8", "utf-8", "errors")

	htmlStr = nodesToString(root)
	fmt.Println(htmlStr)

	p, err := ParagraphObjectModel(htmlStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, para := range p {
		fmt.Printf("%s\n\tText: %s\n\tTagCount: %d\n\tWordCount: %d\n\tLinkedCharCount: %d\n", para.DomPath, para.Text, para.TagCount, para.WordCount, para.LinkedCharCount)
	}

	return htmlStr
}

// nodesToString loops over a node tree and generate HTML string
func nodesToString(node *html.Node) string {
	var response string = ""
	switch node.Type {
		case html.TextNode:
			response = fmt.Sprintf("%s", node.Data)

		case html.ElementNode, html.DocumentNode:
			var att string = ""
			if len(node.Attr)>0 {
				for _, a := range node.Attr {
					att = fmt.Sprintf("%s %s=\"%s\"", att, a.Key, a.Val)
				}
			}

			var content string
			if len(node.Child)>0 {
				for _, n := range node.Child {
					content = fmt.Sprintf("%s%s", content, nodesToString(n))
				}
			}

			response = fmt.Sprintf("<%s%s>%s</%s>", node.Data, att, content, node.Data)

	}
	return response
}
