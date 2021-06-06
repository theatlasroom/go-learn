package main

// TODO: Expose scraped data via graphql
// TODO: scrape paginated data
// TODO: concurrency
import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/theatlasroom/go-learn/lbxd-scrape/lbxd"
)

var films lbxd.Films

const username = "junior1z1337"

func startServer() {
	// Initialize standard Go html template engine
	engine := html.New("./client/dist", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// TODO: check if prod, dont enable
	corsConfig := cors.Config{
		// AllowOrigins: "localhost",
		// AllowHeaders: "Origin, Content-Type, Accept",
	}

	// Or extend your config for customization
	app.Use(cors.New(corsConfig))

	app.Static("/", "./client/dist")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
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
