package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int
	Name string
}

type NFT struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"userID"`
	OwnerName string `json:"owner"`
	Title     string `json:"title"`
}

var nfts = []NFT{
	{1, 1, "cetinboran", "Python Master"},
	{2, 1, "mehmetozler", "C++ Master"},
	{3, 2, "arzutoktas", "C# Master"},
}

func createUsers() []User {
	return []User{
		{ID: 1, Name: "cetinboran"},
		{ID: 2, Name: "kaanmesum"},
	}
}

func main() {
	app := fiber.New()

	users := createUsers()
	currentUser := users[0]

	// VULNERABLE ENDPOINT
	app.Post("/nfts/vuln/get", func(c *fiber.Ctx) error {
		type RequestBody struct {
			NFTID  int `json:"nftid"`
			UserID int `json:"userID"`
		}

		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		for _, nft := range nfts {
			if nft.ID == body.NFTID {
				return c.JSON(nft)
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("NFT not found")
	})

	// SAFE ENDPOINT
	app.Post("/nfts/safe/get", func(c *fiber.Ctx) error {
		type RequestBody struct {
			NFTID  int `json:"nftid"`
			UserID int `json:"userID"`
		}

		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if currentUser.ID != body.UserID {
			return c.Status(fiber.StatusInternalServerError).SendString("IDOR Detected")
		}

		for _, nft := range nfts {
			if nft.ID == body.NFTID {
				return c.JSON(nft)
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("NFT not found or access denied")
	})

	app.Listen(":3000")
}
