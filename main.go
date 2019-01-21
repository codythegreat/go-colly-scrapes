package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("codymaxie.com", "github.io", "lunchscore.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link discovered: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Visits head-fi.org and starts scraping
	c.Visit("https://codymaxie.com")
}
