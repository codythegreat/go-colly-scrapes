package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		fmt.Printf("Quote: %s\nAuthor: %s\n\n", e.ChildText("span.text"), e.ChildText("small.author"))
	})
	c.Visit("http://quotes.toscrape.com/")
}
