// attempt to extract as many artists as possible from Wikipedia

package main

import (
	"fmt"
	"net/http"
	"os"
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

	doc.Find(".div-col.columns.column-width ul li a").Each(func(_ int, s *goquery.Selection) {
		artists = append(artists, genre+" "+s.Text()+"\n")
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

	var results []string

	doc.Find(".div-col.columns.column-width ul li a").Each(func(_ int, s *goquery.Selection) {
		musicType := strings.Replace(s.Text(), "List of ", "", -1)
		link, ok := s.Attr("href")
		if ok {
			results = append(results, getArtistNames(link, musicType)...)

		}
	})

	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(results))
	if err != nil {
		panic(err)
	}
}

func main() {
	getGenreLinks()
}
