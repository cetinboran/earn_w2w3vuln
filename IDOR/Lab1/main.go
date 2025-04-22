package main

// Developed by Çetin Boran Mesüm

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

var usersNfts = []NFT{
	{1, 1, "cetinboran", "Python Master"},
	{2, 1, "cetinboran", "C++ Master"},
	{3, 2, "kaanmesum", "C# Master"},
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

		for _, nft := range usersNfts {
			if nft.OwnerID == body.UserID {
				if nft.ID == body.NFTID {
					return c.JSON(nft)
				}
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

		if body.UserID != 1 && body.UserID != 2 {
			return c.Status(fiber.StatusBadRequest).SendString("There is no such user")
		}

		for _, nft := range usersNfts {
			if nft.OwnerID == body.UserID {
				if nft.ID == body.NFTID {
					if nft.OwnerID != currentUser.ID {
						return c.Status(fiber.StatusInternalServerError).SendString("IDOR Detected")
					}
					return c.JSON(nft)
				}
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("NFT not found or access denied")
	})

	app.Listen(":3000")
}
