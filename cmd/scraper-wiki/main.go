package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("nirmalp.com", "en.wikipedia.org"),
	)

	c.OnHTML("#mw-content-text .mw-parser-output p", func(e *colly.HTMLElement) {

		link := e.ChildAttr("a[href][title]", "href")

		if link == "/wiki/Philosophy" {
			fmt.Println(link)
			os.Exit(1)
		}

		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	err := c.Visit("https://en.wikipedia.org/wiki/Akbar")
	if err != nil {
		fmt.Println(err)
	}
}
