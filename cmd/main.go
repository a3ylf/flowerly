package main

import (
	"fmt"
	"log"

	"github.com/a3ylf/flowerly/database"
	"github.com/gofiber/fiber/v2"
)

// Simula uma base de dados de produtos

func main() {
	app := fiber.New()

	db := database.InitDB()
	// Endpoint para retornar todos os produtos
	app.Get("/plants", func(c *fiber.Ctx) error {
		plants, err := db.GetProducts()
		fmt.Println(plants)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(plants)
	})

	// Servir a página HTML estática
	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}
