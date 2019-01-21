package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.codymaxie.com"),
	)

	c.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting", r.URL.String())
	})
	c.Visit("https://www.codymaxie.com")
}
