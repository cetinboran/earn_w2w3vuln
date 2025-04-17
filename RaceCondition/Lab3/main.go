package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var mu sync.Mutex

type Token struct {
	Name    string
	Balance float64
	Claimed bool
}

var tokens = []Token{
	{Name: "SOL", Balance: 0.2, Claimed: false},
	{Name: "BTC", Balance: 0, Claimed: false},
}

type MyErr struct {
	Error string `json:"error"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/claim/SOL", func(c *fiber.Ctx) error {
		mu.Lock()
		solanaToken := &tokens[0]

		time.Sleep(50 * time.Millisecond)

		if !solanaToken.Claimed {
			// Some Process
			time.Sleep(100 * time.Millisecond)

			solanaToken.Balance++
		} else {
			mu.Unlock()
			return c.Status(400).JSON(MyErr{Error: "Already Claimed!"})
		}
		solanaToken.Claimed = true
		mu.Unlock()

		return c.Status(200).JSON(nil)
	})

	app.Post("/claim/BTC", func(c *fiber.Ctx) error {
		btcToken := &tokens[1]

		time.Sleep(50 * time.Millisecond)

		if !btcToken.Claimed {
			// Some Process
			time.Sleep(100 * time.Millisecond)

			btcToken.Balance++
		} else {
			return c.Status(400).JSON(MyErr{Error: "Already Claimed!"})
		}
		btcToken.Claimed = true

		return c.Status(200).JSON(nil)
	})

	app.Post("/balance/restart", func(c *fiber.Ctx) error {
		tokens = []Token{
			{Name: "SOL", Balance: 0.2, Claimed: false},
			{Name: "BTC", Balance: 0, Claimed: false},
		}

		return c.Status(200).JSON(nil)
	})

	app.Get("/balance", func(c *fiber.Ctx) error {
		response := "Token Balances:\n"
		for _, token := range tokens {
			response += fmt.Sprintf("%s: %.2f\n", token.Name, token.Balance)
		}
		return c.SendString(response)
	})

	app.Listen(":3000")
}
