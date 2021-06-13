package lbxd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// TODO: minimal data from scrape, Collate all the possible links just the ID and url if thats available
// TODO: then LATER do the full scrape once we have collated all the links

const (
	posterContainers  string = ".poster-list .film-poster"
	filmOriginalTitle string = ".frame"
)

// Poster specifies the data for poster
type Poster struct {
	ID     int64
	URL    string
	Height int16
	Width  int16
}

type Posters []Poster

// Film full metadata for a film
type Film struct {
	ID            int64  `json:"id"`
	Name          string `json:"title"`
	OriginalTitle string `json:"searchTitle"`
	Year          int16  `json:"year"`
	URL           string `json:"url"`
	PosterURL     string `json:"posterUrl"`
	ImageURL      string `json:"imageUrl"`
}

type Films []Film

// TODO: type for colly HTMLElement that returns all these helpers
func cleanString(s string) string {
	return strings.Trim(s, " ")
}

func cleanInt(val, errString string, precision int) int64 {
	i, err := strconv.ParseInt(val, 10, precision)

	// TODO: this should probably be custom errors using error.New
	if err != nil {
		errorString := fmt.Errorf(err.Error(), errString, val)
		handleError(errorString)
	}
	return i
}

func cleanInt64(val, errString string) int64 {
	return cleanInt(val, errString, 64)
}

func cleanInt16(val, errString string) int16 {
	return int16(cleanInt(val, errString, 16))
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

func getI16(e *colly.HTMLElement, attr, errString string) int16 {
	val := cleanString(e.Attr(attr))
	return cleanInt16(val, errString)
}

func getI64(e *colly.HTMLElement, attr, errString string) int64 {
	val := cleanString(e.Attr(attr))
	return cleanInt64(val, errString)
}

func getString(e *colly.HTMLElement, attr string) string {
	return cleanString(e.Attr(attr))
}

func getChildString(e *colly.HTMLElement, sel, attr string) string {
	return cleanString(e.ChildAttr(sel, attr))
}

func getImageURL(e *colly.HTMLElement, sel, attr string) string {
	img := e.DOM.Find(sel)
	return cleanString(img.AttrOr(attr, ""))
}

func extractFilm(e *colly.HTMLElement) Film {
	return Film{
		ID:            getI64(e, "data-film-id", "Failed to get ID"),
		Name:          getString(e, "data-film-name"),
		URL:           getString(e, "data-film-slug"),
		PosterURL:     getString(e, "data-poster-url"),
		Year:          getI16(e, "data-film-release-year", "Failed to get Year"),
		OriginalTitle: getChildString(e, filmOriginalTitle, "title"),
		// TODO: might need to do extra extraction for metadata, or maybe we resolve the image later from another source
		ImageURL: getImageURL(e, "image.img", "src"),
	}
}

func extractPoster(e *colly.HTMLElement) Poster {
	h := cleanInt16(getChildString(e, ".image", "height"), "Failed to get height")
	w := cleanInt16(getChildString(e, ".image", "width"), "Failed to get width")
	return Poster{
		ID:     getI64(e, "data-film-id", "Failed to get ID"),
		URL:    getString(e, "data-film-slug"),
		Height: h,
		Width:  w,
	}
}

func (p Poster) metadataURLString() string {
	return fmt.Sprintf("https://letterboxd.com/ajax/poster%vmenu/linked/%dx%d/", p.URL, p.Width, p.Height)
}

func (p Poster) metadataURL() *url.URL {
	u, err := url.Parse(p.metadataURLString())
	handleError(err)

	return u
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init() *colly.Collector {
	c := colly.NewCollector()
	return c
}

func watchlistURL(username string) string {
	return fmt.Sprintf("https://letterboxd.com/%v/watchlist/", username)
}

// FetchFilmsWatchList takes a letterboxd username and fetches the films from their watch list
func FetchFilmsWatchList(username string, c colly.Collector) Films {
	var ps Posters
	var fs Films

	target := watchlistURL(username)

	// Extract all visible movie titles
	c.OnHTML(posterContainers, func(e *colly.HTMLElement) {
		nf := extractPoster(e)
		target := nf.metadataURLString()

		visited, err := c.HasVisited(target)
		handleError(err)

		if !visited {
			err := c.Request(http.MethodGet, target, nil, nil, nil)
			handleError(err)
			ps = append(ps, nf)
		}
	})
	// ajax response
	c.OnHTML(".linked-film-poster", func(e *colly.HTMLElement) {
		fs = append(fs, extractFilm(e))
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("Finished scraping", r.Request.URL)
	// })

	// Kick off the scraper and block until we complete
	c.Visit(target)

	c.Wait()
	return fs
}
