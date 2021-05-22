package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

const target = "https://letterboxd.com/junior1z1337/watchlist/"

// TODO: minimal data from scrape, Collate all the possible links just the ID and url if thats available
// TODO: then LATER do the full scrape once we have collated all the links

const (
	filmContainers string = ".poster-list .linked-film-poster"
	// filmContainers    string = ".poster-list .film-poster.linked-film-poster"
	filmOriginalTitle string = "a.frame"
)

type film struct {
	ID   int64
	Name string
	// used to query for the specific movie
	OriginalTitle string
	Year          int8
	URL           string
	PosterURL     string
}

type films []film

// TODO: type for colly HTMLElement that returns all these helpers
func cleanString(s string) string {
	return strings.Trim(s, " ")
}

func cleanInt(val, errString string, precision int) int64 {
	i, err := strconv.ParseInt(val, 10, precision)
	if err != nil {
		str := fmt.Errorf(err.Error(), errString, val)
		log.Fatal(str)
	}
	return i
}

func cleanInt64(val, errString string) int64 {
	return cleanInt(val, errString, 64)
}

func cleanInt8(val, errString string) int8 {
	return int8(cleanInt(val, errString, 8))
}

// TODO: use this as the struct for the getters
type collyAttr struct {
	attribute   string
	errorString string
}

func getI8(e *colly.HTMLElement, attr, errString string) int8 {
	val := cleanString(e.Attr(attr))
	return cleanInt8(val, errString)
}

func getI64(e *colly.HTMLElement, attr, errString string) int64 {
	val := cleanString(e.Attr(attr))
	return cleanInt64(val, errString)
}

func getString(e *colly.HTMLElement, attr string) string {
	return cleanString(e.Attr(attr))
}

func getChildString(e *colly.HTMLElement, sel, attr string) string {
	// fmt.Println("OT", e.ChildAttr(sel, attr), e)
	el := e.DOM.Find(sel)
	fmt.Println(e.Name, e.Text, e.Attr("class"))
	v, ok := el.Attr(attr)
	if !ok {
		fmt.Println("BUMMER")
	}
	fmt.Println("OT", e.ChildAttr(sel, attr), v)
	// fmt.Println("OT", e.ChildAttr(sel, attr), e.DOM.Find(sel).Attr(attr))
	return cleanString(e.ChildAttr(sel, attr))
}

// func getChildText(e *colly.HTMLElement, sel string) string {
// 	// fmt.Println(e.ChildText(sel))
// 	ed := e.DOM.Closest(sel)
// 	fmt.Println(ed.Attr())
// 	return cleanString(ed.Text())
// 	// return cleanString(e.ChildText(sel))
// }

func extractFilm(e *colly.HTMLElement) film {
	// fmt.Println(e.DOM.Html())
	return film{
		ID:        getI64(e, "data-film-id", "Failed to get ID"),
		Name:      getString(e, "data-film-name"),
		URL:       getString(e, "data-target-link"),
		PosterURL: getString(e, "data-poster-url"),
		// Year:          getI8(e, "data-film-release-year", "Failed to get Year"),
		OriginalTitle: getChildString(e, filmOriginalTitle, "data-original-title"),
	}
}

func main() {
	fmt.Println("Scraping", target)
	c := colly.NewCollector()

	var fs films

	// Extract all visible movie titles
	c.OnHTML(filmContainers, func(e *colly.HTMLElement) {
		nf := extractFilm(e)
		fmt.Println(nf)
		fs = append(fs, nf)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping", r.Request.URL)
	})

	// Move through the pagination
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(target)
}
