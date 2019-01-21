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
		tagString := ""
		e.ForEach("a.tag", func(_ int, elem *colly.HTMLElement) {
			tagString += elem.Text + " "
})
		fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s \n", e.ChildText("span.text"), e.ChildText("small.author"), tagString)
	})
	c.Visit("http://quotes.toscrape.com/")
}
