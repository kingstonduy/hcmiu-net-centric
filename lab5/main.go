package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Serve static assets (images, CSS, JS files, etc.)
	app.Static("/assets", "./assets") // Serve files from the "assets" directory

	app.Get("/exc1", func(c *fiber.Ctx) error {
		str := c.Body()
		return c.SendString(string(str))
	})

	app.Get("/exc2", func(c *fiber.Ctx) error {
		htmlResponse := "<html>Welcome to Very Simple Web Server</html>"
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(htmlResponse)
	})

	// Serve index.html file on /exc3 route
	app.Get("/exc3", func(c *fiber.Ctx) error {
		fileContent, err := os.ReadFile("./exc3.html") // Reads the content of index.html
		if err != nil {
			log.Println("Error reading file:", err)
			return c.Status(500).SendString("Internal Server Error")
		}
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(string(fileContent)) // Sends the file content as response
	})

	app.Get("/exc4", func(c *fiber.Ctx) error {
		fileContent, err := os.ReadFile("./exc4.html") // Reads the content of index.html
		if err != nil {
			log.Println("Error reading file:", err)
			return c.Status(500).SendString("Internal Server Error")
		}
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(string(fileContent)) // Sends the file content as response
	})

	app.Listen(":3000")
}
