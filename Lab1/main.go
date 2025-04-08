package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type NFT struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
}

var nfts = []NFT{
	{"1", "cetinboran", "Python Master"},
	{"2", "mehmetozler", "C++ Master"},
	{"3", "arzutoktas", "C# Master"},
}

func main() {
	app := fiber.New()

	currentUser := "user1"
	fmt.Println(currentUser)

	app.Get("/nfts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, nft := range nfts {
			if nft.ID == id {
				return c.JSON(nft)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("NFT not found")
	})

	app.Listen(":3000")
}
