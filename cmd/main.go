package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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
	app.Get("/vendors", func(c *fiber.Ctx) error {
		vendors, err := db.GetVendors()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to fetch vendors")
		}
		return c.JSON(vendors)
	})
	app.Get("/cookies", func(c *fiber.Ctx) error {
		ret := ""
		current := c.Cookies("vendor")
		if current == "" {
			ret = fmt.Sprint(ret + "\nCookie para Vendor não encontrado")
		} else {
			ret = fmt.Sprint(ret + "\nValor do cookie Vendor: " + current)
		}
		current = c.Cookies("client")
		if current == "" {
			ret = fmt.Sprint(ret + "\nCookie para cliente não encontrado")
		} else {
			ret = fmt.Sprint(ret + "\nValor do cookie Cliente: " + current)
		}
		return c.Status(fiber.StatusOK).SendString(ret)
	})

}
func setupRoutes(app *fiber.App, db *database.Database) {
	type login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	app.Get("/signup/client", func(c *fiber.Ctx) error {
		return c.Render("signupClient", fiber.Map{}) // Serve o arquivo HTML
	})

	app.Post("/signup/client", func(c *fiber.Ctx) error {
		client := new(database.Client)
		if err := c.BodyParser(client); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
		}
		// Consulta SQL para inserir o cliente
		query := `INSERT INTO client (name, email, password, cpf, rua, num) 
		          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

		err := auth.EnsureSignup(&client.Vendor)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		password, err := auth.HashPassword(client.Password)
		_, err = db.Create(query, client.Name, client.Email, password, client.CPF, client.Rua, client.Num)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to create client: %v", err))
		}

		return c.SendString("Client created successfully")

	},
	)
	app.Get("/login/vendor", func(c *fiber.Ctx) error {
		return c.Render("loginVendor", fiber.Map{}) // Serve o arquivo HTML
	})
	app.Get("/login/client", func(c *fiber.Ctx) error {
		return c.Render("loginClient", fiber.Map{}) // Serve o arquivo HTML
	})
	app.Post("/login/client", func(c *fiber.Ctx) error {
		login := new(login)
		if err := c.BodyParser(login); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
		}
		log.Println(login.Email + " se conectou")
		name, psw, err := db.GetLogin("client", login.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())

		}

		if err = auth.CheckPassword([]byte(psw), []byte(login.Password)); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Wrong password")

		}
		cookie := new(fiber.Cookie)
		cookie.Name = "client"
		cookie.Value = name
		cookie.Expires = time.Now().Add(3 * time.Hour)
		c.Cookie(cookie)
		return c.SendString(fmt.Sprintf("Login feito corretamente para cliente de nome: %s", name))

	})
	app.Post("/login/vendor", func(c *fiber.Ctx) error {
		login := new(login)
		if err := c.BodyParser(login); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
		}
		fmt.Println(login.Email)
		fmt.Println(login.Password)
		name, psw, err := db.GetLogin("vendor", login.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())

		}

		if err = auth.CheckPassword([]byte(psw), []byte(login.Password)); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Wrong password")

		}
		cookie := new(fiber.Cookie)
		cookie.Name = "vendor"
		cookie.Value = name
		cookie.Expires = time.Now().Add(3 * time.Hour)
		c.Cookie(cookie)
		return c.SendString(fmt.Sprintf("Login feito corretamente para vendedor de nome: %s", name))

	})
	app.Post("/signup/vendor", func(c *fiber.Ctx) error {
		client := new(database.Client)
		if err := c.BodyParser(client); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
		}
		// Consulta SQL para inserir o vendedor
		query := `INSERT INTO vendor (name, email, password, cpf) 
		          VALUES ($1, $2, $3, $4)`

		err := auth.EnsureSignup(&client.Vendor)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		password, err := auth.HashPassword(client.Password)
		_, err = db.Create(query, client.Name, client.Email, password, client.CPF, client.Rua, client.Num)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to create vendor: %v", err))
		}

		return c.SendString("Vendor created successfully")

	},
	)
	app.Get("/logout", func(c *fiber.Ctx) error {

		cookie := c.Cookies("client")

		if cookie != "" {
			c.Cookie(&fiber.Cookie{
				Name:    "client",
				Expires: time.Now().Add(-time.Hour),
			})
		}
		cookie = c.Cookies("vendor")

		if cookie != "" {
			c.Cookie(&fiber.Cookie{
				Name:    "vendor",
				Expires: time.Now().Add(-time.Hour),
			})
		}

		return c.SendString("Todos os cookies foram deletados!")
	})

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
		return c.JSON(fiber.Map{
			"Title":  "Todas as plantas a venda",
			"Plants": plants,
		})
	})
	app.Get("/plants/mari", func(c *fiber.Ctx) error {
		plants, err := db.GetProductsFromMari()
		if err != nil {
			return err
		}
		return c.JSON( fiber.Map{
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
		return c.JSON(fiber.Map{
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
		return c.JSON(fiber.Map{
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
			if err.Error() == fmt.Sprintf("Nenhuma planta encontrada com o nome; %s", name) {
				return c.Status(fiber.StatusNotFound).SendString("Couldn't find plant named: " + name)
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching plant")
		}

		return c.JSON(fiber.Map{
			"Title": "Planta: " + name,
			"Plant": plant,
		})
	})

}
