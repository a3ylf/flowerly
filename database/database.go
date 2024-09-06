package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Database struct {
	db *sql.DB
}

func InitDB() *Database {
	// Carregar o .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", dbUser, dbPassword, dbName)

	// Abrir conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Verificar se a conexão está funcional
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	return &Database{
		db: db,
	}
}

func (db *Database) GetProducts() ([]Plant, error) {
	rows, err := db.db.Query("SELECT * FROM plants")
	if err != nil {
		return []Plant{}, err
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.ID, &plant.Name, &plant.ScientificName, &plant.Description, &plant.Category, &plant.Price, &plant.StockQuantity, &plant.ImageURL, &plant.OriginLocation); err != nil {
			return []Plant{}, err
		}
		plants = append(plants, plant)
	}

	// Verificar por erros que possam ter ocorrido após a iteração
	if err = rows.Err(); err != nil {
		return []Plant{}, err
	}

	return plants, nil
}
func (db *Database) GetProductByName(name string) (Plant, error) {
	rows := db.db.QueryRow("SELECT * FROM plants WHERE name = $1 LIMIT 1",name)

	var plant Plant
	if err := rows.Scan(&plant.ID, &plant.Name, &plant.ScientificName, &plant.Description, &plant.Category, &plant.Price, &plant.StockQuantity, &plant.ImageURL, &plant.OriginLocation); err != nil {
	    if err == sql.ErrNoRows {
            return Plant{},fmt.Errorf("Nenhuma flor encontrada com o nome; %s", name)
        }
		return Plant{}, err
	}
	return plant, nil
}
func (db *Database) GetProductsFromMari() ([]Plant, error) {
	rows, err := db.db.Query("SELECT * FROM plants WHERE origin_location = $1","Mari")
	if err != nil {
		return []Plant{}, err
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.ID, &plant.Name, &plant.ScientificName, &plant.Description, &plant.Category, &plant.Price, &plant.StockQuantity, &plant.ImageURL, &plant.OriginLocation); err != nil {
			return []Plant{}, err
		}
		plants = append(plants, plant)
	}

	// Verificar por erros que possam ter ocorrido após a iteração
	if err = rows.Err(); err != nil {
		return []Plant{}, err
	}

	return plants, nil
}


func (db *Database) GetProductsByCategory(category string) ([]Plant, error) {
    category = strings.ToLower(category)
	rows, err := db.db.Query("SELECT * FROM plants WHERE category = $1",category)
	if err != nil {
		return []Plant{}, err
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.ID, &plant.Name, &plant.ScientificName, &plant.Description, &plant.Category, &plant.Price, &plant.StockQuantity, &plant.ImageURL, &plant.OriginLocation); err != nil {
			return []Plant{}, err
		}
		plants = append(plants, plant)
	}

	// Verificar por erros que possam ter ocorrido após a iteração
	if err = rows.Err(); err != nil {
		return []Plant{}, err
	}

	return plants, nil
}
func (db *Database) GetProductsByPrice(price int) ([]Plant, error) {
	rows, err := db.db.Query("SELECT * FROM plants WHERE price <= $1",price)
	if err != nil {
		return []Plant{}, err
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.ID, &plant.Name, &plant.ScientificName, &plant.Description, &plant.Category, &plant.Price, &plant.StockQuantity, &plant.ImageURL, &plant.OriginLocation); err != nil {
			return []Plant{}, err
		}
		plants = append(plants, plant)
	}

	// Verificar por erros que possam ter ocorrido após a iteração
	if err = rows.Err(); err != nil {
		return []Plant{}, err
	}

	return plants, nil
}
