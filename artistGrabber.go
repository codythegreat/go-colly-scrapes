// attempt to extract as many artists as possible from Wikipedia and output to json

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// structure that artist data is output as
type Artist struct {
	Name  string
	Genre string
	Id    string `json:"id"`
}

var artistData []Artist

const genreList string = "https://en.wikipedia.org/wiki/Lists_of_musicians?oldformat=true"

// hold all links found in genreList
var genreListLinks = make([]string, 0)

// start at 1, iterate 1 for each artist
var currentID int64 = 1

// get all artists in a genre's list of artists
func getArtistNames(link, genre string) []string {
	// make an empty slice to store artists
	artists := make([]string, 0)
	// attempt to fetch webpage
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
	// drill down to element containing artist name
	doc.Find(".div-col.columns.column-width ul li a").Each(func(_ int, s *goquery.Selection) {
		// test each result with regex to make sure it is an artist
		match, err := regexp.MatchString(`\[\d+\]`, s.Text())
		if err != nil {
			panic(err)
		}
		matchList, err := regexp.MatchString(`^[Ll]ists?.*`, s.Text())
		if err != nil {
			panic(err)
		}
		// if neither regex matched, create new Artist and add to artistData slice
		if match == false && matchList == false {
			artist := Artist{s.Text(), genre, strconv.FormatInt(currentID, 10)}
			currentID++
			artistData = append(artistData, artist)
			artists = append(artists, genre+" "+s.Text()+"\n") // old way of printing output. leaving for reference
		}
	})
	return artists
}

func getGenreLinks() {
	// attempt to fetch webpage
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
	// for each link in genreList, get artist names from that link
	doc.Find(".div-col.columns.column-width ul li a").Each(func(_ int, s *goquery.Selection) {
		musicType := strings.Replace(s.Text(), "List of ", "", -1)
		link, ok := s.Attr("href")
		if ok {
			results = append(results, getArtistNames(link, musicType)...)

		}
	})
	// create output file (used before implementing json format
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	// write scrapped data to output file
	_, err = file.WriteString(fmt.Sprintln(results))
	if err != nil {
		panic(err)
	}

	// create []byte var to store json data
	var jsonData []byte
	// pretty format the data with MarshalIndent
	jsonData, err = json.MarshalIndent(artistData, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	// print the data to console, and out put to json file with the following:
	// go run artistGrabber.go >> path/to/output.json
	fmt.Println(string(jsonData))
}

func main() {
	getGenreLinks()
}
