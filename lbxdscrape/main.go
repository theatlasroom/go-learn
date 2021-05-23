package main

// TODO: Expose scraped data via graphql
// TODO: scrape paginated data
// TODO: concurrency

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/theatlasroom/go-learn/lbxd-scrape/lbxd"
)

const username = "junior1z1337"

func main() {
	fmt.Println("Scraping", username)
	c := *colly.NewCollector()
	films := lbxd.FetchFilmsWatchList(username, c)
	c.Wait()

	fmt.Println("DONESKIES!", films)
}
