package main

import (
	"encoding/json"
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

	// Criando o app Fiber e configurando a engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	db := database.InitDB()
	setupTestRoutes(app, db)
	setupRoutes(app, db)

	// Servir a página HTML estática
	app.Static("/", "./views")

	log.Fatal(app.Listen(":3000"))
}

type logincookie struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type item struct {
	Id      int `json:"id"`
	Ammount int `json:"ammount"`
}
type cartcookie struct {
	Itens []item `json:"itens"`
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
		cook := new(logincookie)
		var err error
		ret := ""
		current := c.Cookies("vendor")

		if current == "" {
			ret = fmt.Sprint(ret + "\nCookie para Vendor não encontrado")
		} else {
			err := json.Unmarshal([]byte(current), cook)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Error Unmarshalling")
			}
			ret = fmt.Sprint(ret+"\nValor do cookie Vendor: "+cook.Name+" Id: ", cook.Id)
		}
		current = c.Cookies("client")
		if current == "" {
			ret = fmt.Sprint(ret + "\nCookie para cliente não encontrado")
		} else {
			err = json.Unmarshal([]byte(current), cook)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Error Unmarshalling")
			}

			ret = fmt.Sprint(ret + "\nValor do cookie Cliente: " + current)
		}
		cartcook := new(cartcookie)
		current = c.Cookies("cart")
		if current == "" {
			ret = fmt.Sprint(ret + "\nCookie cart não encontrado")
		} else {
			err = json.Unmarshal([]byte(current), cartcook)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Error Unmarshalling")
			}

			ret = fmt.Sprint(ret + "\nValor do cookie Cart: " + current)
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
	app.Get("/addplant/:id/:ammount", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to fetch params: ", err))
		}
		ammount, err := c.ParamsInt("ammount")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to fetch params: ", err))
		}

		cartcontent := c.Cookies("cart")
		if cartcontent == "" {
			jcart, err := json.Marshal(cartcookie{
				Itens: []item{
					item{Id: id,
						Ammount: ammount,
					},
				},
			})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error marshaling: ", err))
			}
			cookie := new(fiber.Cookie)
			cookie.Name = "cart"
			cookie.Value = string(jcart)
			cookie.Expires = time.Now().Add(3 * time.Hour)
			cookie.HTTPOnly = false
			c.Cookie(cookie)
			return c.Status(fiber.StatusOK).SendString("Succesfully created a new cart and added it to it")
		}
		cart := new(cartcookie)
		err = json.Unmarshal([]byte(cartcontent), cart)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error unmarshaling: ", err))
		}
		already := false
		for item := range cart.Itens {
			if cart.Itens[item].Id == id {
				cart.Itens[item].Ammount += ammount
				already = true
				break
			}
		}
		if !already {
			cart.Itens = append(cart.Itens, item{Id: id, Ammount: ammount})
		}

		newcookie, err := json.Marshal(cart)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error marshaling: ", err))
		}
		cookie := new(fiber.Cookie)
		cookie.Name = "cart"
		cookie.Value = string(newcookie)
		cookie.Expires = time.Now().Add(3 * time.Hour)
		cookie.HTTPOnly = false
		c.Cookie(cookie)

		return c.Status(fiber.StatusOK).SendString("Succesfully put it in the cart")

	},
	)

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
		id, name, psw, err := db.GetLogin("client", login.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())

		}

		if err = auth.CheckPassword([]byte(psw), []byte(login.Password)); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Wrong password")

		}
		clientcookie := logincookie{
			Id:   id,
			Name: name,
		}
		value, err := json.Marshal(clientcookie)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error marshaling json")
		}
		cookie := new(fiber.Cookie)
		cookie.Name = "client"
		cookie.Value = string(value)
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
		id, name, psw, err := db.GetLogin("vendor", login.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())

		}
		clientcookie := logincookie{
			Id:   id,
			Name: name,
		}
		value, err := json.Marshal(clientcookie)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error marshaling json")
		}

		if err = auth.CheckPassword([]byte(psw), []byte(login.Password)); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Wrong password")

		}
		log.Println(login.Email + " se conectou")
		cookie := new(fiber.Cookie)
		cookie.Name = "vendor"
		cookie.Value = string(value)
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

		c.ClearCookie()

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
		return c.JSON(fiber.Map{
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
