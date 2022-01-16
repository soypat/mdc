package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// icons.go auto-generator.
func main() {
	var rd io.Reader

	rd, err := os.Open("icons.html")
	if err != nil {
		fmt.Println("No icons.html")
	}
	z := html.NewTokenizer(rd)
	count := 0
	depth := 0
	currentTag := ""
	var lab [2]string
	labelsFound := make(map[string]struct{})
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			fmt.Printf("\n\nEND PROGRAM")
			os.Exit(0)
		case html.TextToken:
			if currentTag != "span" && depth == 0 {
				continue
			}
			str := string(z.Text())
			// fmt.Println(str)
			lab[count] = strings.TrimSpace(str)
			count++
			if count >= 2 {
				count = 0
				if _, pr := labelsFound[lab[0]]; pr {
					continue
				}
				labelsFound[lab[0]] = struct{}{}
				fmt.Printf("%s Icon = %q\n", strings.Join(strings.Fields(lab[1]), ""), lab[0])
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
			currentTag = string(tn)
		}
	}
}
