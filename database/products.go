package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

func (db *Database) GetProducts() ([]Plant, error) {
	rows, err := db.Db.Query("SELECT * FROM plants")
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
	rows := db.Db.QueryRow("SELECT * FROM plants WHERE name = $1 LIMIT 1", name)

	var plant Plant
	if err := rows.Scan(&plant.ID, &plant.Name, &plant.ScientificName, &plant.Description, &plant.Category, &plant.Price, &plant.StockQuantity, &plant.ImageURL, &plant.OriginLocation); err != nil {
		if err == sql.ErrNoRows {
			return Plant{}, fmt.Errorf("Nenhuma flor encontrada com o nome; %s", name)
		}
		return Plant{}, err
	}
	return plant, nil
}
func (db *Database) GetProductsFromMari() ([]Plant, error) {
	rows, err := db.Db.Query("SELECT * FROM plants WHERE origin_location = $1", "Mari")
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
	rows, err := db.Db.Query("SELECT * FROM plants WHERE category = $1", category)
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
	rows, err := db.Db.Query("SELECT * FROM plants WHERE price <= $1", price)
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
func(db *Database) CreatePurchase(cartID int, totalAmount float64, paymentStatus, paymentMethod string) (int, error) {
    // Query para inserir uma nova compra
    query := `
        INSERT INTO purchase (cart_id, total_amount, payment_status, payment_method)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

    var purchaseID int

    // Executando a query de inserção e retornando o ID da nova compra
    err := db.Db.QueryRow(query, cartID, totalAmount, paymentStatus, paymentMethod).Scan(&purchaseID)
    if err != nil {
        return 0, fmt.Errorf("erro ao inserir compra: %v", err)
    }

    return purchaseID, nil
}
func(db *Database) GetClientPurchases(clientID int) (*ClientPurchases, error) {
    query := `
        SELECT client_id, client_name, purchase_id, purchase_date, total_amount, 
               payment_status, payment_method, product_id, product_name, quantity
        FROM client_purchases_view
        WHERE client_id = $1
        ORDER BY purchase_id
    `

    rows, err := db.Db.Query(query, clientID)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar pedidos: %v", err)
    }
    defer rows.Close()

    var clientPurchases ClientPurchases
    var lastPurchaseID int
    var currentPurchase *Purchase

    for rows.Next() {
        var product Product
        var purchase Purchase

        err := rows.Scan(&clientPurchases.ClientID, &clientPurchases.ClientName, &purchase.PurchaseID,
            &purchase.PurchaseDate, &purchase.TotalAmount, &purchase.PaymentStatus,
            &purchase.PaymentMethod, &product.ProductID, &product.ProductName, &product.Quantity)
        if err != nil {
            return nil, fmt.Errorf("erro ao escanear resultado: %v", err)
        }

        // Verifica se é uma nova compra
        if purchase.PurchaseID != lastPurchaseID {
            // Adiciona a compra anterior (se existir)
            if currentPurchase != nil {
                clientPurchases.Purchases = append(clientPurchases.Purchases, *currentPurchase)
            }
            // Inicia uma nova compra
            currentPurchase = &purchase
            currentPurchase.Products = []Product{}
            lastPurchaseID = purchase.PurchaseID
        }

        // Adiciona o produto à compra atual
        currentPurchase.Products = append(currentPurchase.Products, product)
    }

    // Adiciona a última compra
    if currentPurchase != nil {
        clientPurchases.Purchases = append(clientPurchases.Purchases, *currentPurchase)
    }

    return &clientPurchases, nil
}