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
	setupRoutes(app, db)
	// Endpoint para retornar todos os produtos

	// Servir a página HTML estática
	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}
func setupRoutes(app *fiber.App, db *database.Database) {
	app.Get("/plants", func(c *fiber.Ctx) error {
		plants, err := db.GetProducts()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(plants)
	})
    app.Get("/plants/:category", func(c *fiber.Ctx) error {
        category := c.Params("category")
		plants, err := db.GetProductsByCategory(category)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(plants)
	})

	// Rota que pega o nome diretamente no caminho
	app.Get("/plant/:name", func(c *fiber.Ctx) error {
		name := c.Params("name") // Pega o nome do parâmetro de caminho

		plant, err := db.GetProductByName(name)
		if err != nil {
			if err.Error() == fmt.Sprintf("Nenhuma flor encontrada com o nome; %s", name) {
				return c.Status(fiber.StatusNotFound).SendString("Couldn't find plant named: " + name)
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching plant")
		}

		return c.JSON(plant)
	})

}
