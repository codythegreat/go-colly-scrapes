// gathers a list of composers from wikipedia
package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_composers_by_name")
	if err != nil {
		fmt.Println("Error while retrieving page")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("error creating goquery document")
	}

	doc.Find(".div-col.columns.column-width").Find("ul").Each(func(_ int, s *goquery.Selection) {
		s.Find("li").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("Composer: %s\n", s.Text())
		})
	})
}
