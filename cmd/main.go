package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/a3ylf/flowerly/auth"
	"github.com/a3ylf/flowerly/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Simula uma base de dados de produtos

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db := database.InitDB()
	setupTestRoutes(app, db)
	setupRoutes(app, db)
	// Endpoint para retornar todos os produtos

	// Servir a página HTML estática
	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}

func setupTestRoutes(app *fiber.App, db *database.Database) {
	app.Get("/clients", func(c *fiber.Ctx) error {
		clients, err := db.GetClients()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to fetch clients")
		}
		return c.JSON(clients)
	})

}
func setupRoutes(app *fiber.App, db *database.Database) {

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{}) // Serve o arquivo HTML
	})

	app.Post("/signup/client", func(c *fiber.Ctx) error {
		client := new(database.Client)
		if err := c.BodyParser(client); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
		}
		// Consulta SQL para inserir o cliente
		query := `INSERT INTO client (name, email, password, cpf, rua, num) 
		          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

		password, err := auth.HashPassword(client.Password)
		id, err := db.Create(query, client.Name, client.Email, password, client.CPF, client.Rua, client.Num)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to create client: %v", err))
		}

		return c.SendString(fmt.Sprintf("Client created with ID: %d", id))

	},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "FLOWERLY",
		})
	})

	app.Get("/plants/all", func(c *fiber.Ctx) error {
		plants, err := db.GetProducts()
		if err != nil {
			return err
		}
		return c.Render("view-plants", fiber.Map{
			"Title":  "Todas as plantas a venda",
			"Plants": plants,
		})
	})
	app.Get("/plants/mari", func(c *fiber.Ctx) error {
		plants, err := db.GetProductsFromMari()
		if err != nil {
			return err
		}
		return c.Render("view-plants", fiber.Map{
			"Title":  "Todas as plantas a venda de mari (Imperdiveis)",
			"Plants": plants,
		})
	})

	app.Get("/plants/category/", func(c *fiber.Ctx) error {
		category := c.Query("category")
		plants, err := db.GetProductsByCategory(category)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Render("view-plants", fiber.Map{
			"Title":  "Todas as plantas da categoria " + category,
			"Plants": plants,
		})
	})
	app.Get("/plants/price/", func(c *fiber.Ctx) error {
		max := c.Query("max")
		if max == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Coloque um valor!",
			})
		}
		maximus, err := strconv.Atoi(max)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Não pode ser letra né ",
			})
		}
		plants, err := db.GetProductsByPrice(maximus)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Render("view-plants", fiber.Map{
			"Title":  "Todas as plantas de valor abaixo de " + max,
			"Plants": plants,
		})
	})
	// Rota que pega o nome diretamente no caminho
	app.Get("/plant/name/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		name = strings.NewReplacer("%20", " ").Replace(name)
		plant, err := db.GetProductByName(name)

		if err != nil {
			if err.Error() == fmt.Sprintf("Nenhuma flor encontrada com o nome; %s", name) {
				return c.Status(fiber.StatusNotFound).SendString("Couldn't find plant named: " + name)
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching plant")
		}

		return c.Render("view-full-plant", fiber.Map{
			"Title": "Planta: " + name,
			"Plant": plant,
		})
	})

}
