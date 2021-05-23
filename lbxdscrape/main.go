package main

// TODO: Expose scraped data via graphql
// TODO: scrape paginated data
// TODO: concurrency
import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/theatlasroom/go-learn/lbxd-scrape/lbxd"
)

var films lbxd.Films

const username = "junior1z1337"

func startServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/films", func(c *fiber.Ctx) error {
		return c.JSON(films)
	})

	app.Listen(":3000")
}

func main() {
	fmt.Print(fmt.Sprintf("Preparing data for %v %v ", username, "..."))
	c := *colly.NewCollector()
	films = lbxd.FetchFilmsWatchList(username, c)
	c.Wait()
	fmt.Print("Done")

	startServer()
}
