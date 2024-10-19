package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Database struct {
	Db *sql.DB
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
		Db: db,
	}
}
func (db Database) GetLogin(table, email string) (int, string, error) {
	var id int
	var psw string
<<<<<<< Updated upstream
	err := db.db.QueryRow(fmt.Sprintf("SELECT id, password FROM %s WHERE email = $1;", table), email).Scan(&id, &psw)
=======
	err := db.Db.QueryRow(fmt.Sprintf("SELECT id, name, password FROM %s WHERE email = $1;", table), email).Scan(&id, &name, &psw)
>>>>>>> Stashed changes
	if err != nil {

		return 0, "", err
	}
	return id, psw, nil
}

func (db Database) Create(query string, args ...interface{}) (int, error) {
	var id int
	err := db.Db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create record: %v", err)
	}
	return id, nil
}
func (db Database) Createcartitem(query string, args ...interface{}) (error) {
	_,err := db.Db.Exec(query, args...)
	if err != nil {
		return  fmt.Errorf("failed to create record: %v", err)
	}
	return  nil
}