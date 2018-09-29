package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is xl94.zhang@gmail.com
email1 is abc.d@def.org
emial2 is kkk.g@qqq.com
`

func main() {
	re := regexp.MustCompile(
		`[a-zA-Z0-9]+\.[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}
