package main

import (
	"Task_spider_web_go/resources"
	"fmt"
	"net/url"
	"os"
)

func main() {
	var link string
	fmt.Print("Введите ссылку:")
	fmt.Fscan(os.Stdin, &link)
	url, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	resources.Reading(url.Scheme, url.Hostname())
}
