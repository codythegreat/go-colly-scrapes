// attempt to extract as many artists as possible from Wikipedia

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const genreList string = "https://en.wikipedia.org/wiki/Lists_of_musicians?oldformat=true"

var genreListLinks = make([]string, 0)

func getArtistNames(link, genre string) []string {
	artists := make([]string, 0)

	res, err := http.Get("https://en.wikipedia.org" + link)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %s", link)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find(".div-col.columns.column-width").Find("ul").Find("li").Find("a").Each(func(_ int, s *goquery.Selection) {
		artists = append(artists, genre+" "+s.Text())
	})
	return artists
}

func getGenreLinks() {
	res, err := http.Get(genreList)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find(".div-col.columns.column-width").Find("ul").Find("li").Find("a").Each(func(_ int, s *goquery.Selection) {
		musicType := strings.Replace(s.Text(), "List of ", "", -1)
		fmt.Printf("%s\n", musicType)
		link, ok := s.Attr("href")
		if ok {
			fmt.Println(strings.Join(getArtistNames(link, musicType), "\n"))
		}
	})
}

func main() {
	getGenreLinks()
}
