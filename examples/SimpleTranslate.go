// +build ignore

package main

import (
	"fmt"
	"github.com/Conight/go-googletrans"
)

var content = `你好，世界！`

func main() {
	c := translate.Config{
		Proxy: "http://127.0.0.1:1087",
		UserAgent: []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	t := translate.New(c)
	result, err := t.Translate(content, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
