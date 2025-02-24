package database

import (
	"errors"
)

func (db *Database) GetClients() ([]Client, error) {
	rows, err := db.Db.Query(`SELECT id, name, email, password, cpf, rua, num FROM client`)
	if err != nil {
		return []Client{}, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Password, &client.CPF, &client.Rua, &client.Num); err != nil {
			return []Client{}, err
		}
		clients = append(clients, client)
	}
	return clients, nil

}
func (db *Database) GetVendors() ([]Vendor, error) {
	rows, err := db.Db.Query(`SELECT id, name, email, password, cpf FROM vendor`)
	if err != nil {
		return []Vendor{}, err
	}
	defer rows.Close()

	var vendors []Vendor
	for rows.Next() {
		var vendor Vendor
		if err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Email, &vendor.Password, &vendor.CPF); err != nil {
			return []Vendor{}, err
		}
		vendors = append(vendors, vendor)
	}
	return vendors, nil

}
func (db *Database) GetClient(id int) (Client, error) {
	rows, err := db.Db.Query(`SELECT id, name, email, password, cpf, rua, num FROM client where id = $1`, id)
	if err != nil {
		return Client{}, err
	}
	defer rows.Close()
	client := Client{}
	if rows.Next() {

		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Password, &client.CPF, &client.Rua, &client.Num); err != nil {
			return Client{}, err
		}
	}
	return client, nil

}
func (db *Database) GetVendor(id int) (Vendor, error) {
	rows, err := db.Db.Query(`SELECT id, name, email, password, cpf FROM vendor where id = $1`, id)
	if err != nil {
		return Vendor{}, err
	}
	defer rows.Close()

	var vendor Vendor
	if rows.Next() {
		if err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Email, &vendor.Password, &vendor.CPF); err != nil {
			return Vendor{}, err
		}
	}
	return vendor, nil

}

// db.go
func (db *Database) GetPendingPurchases() ([]ClientPurchases, error) {
	query := `
        SELECT client_id, 
               purchase_id, 
               purchase_date, 
               total_amount, 
               payment_status, 
               payment_method 
        FROM client_purchases_view
        WHERE payment_status = 'esperando efetivação'
    `

	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, errors.New("erro ao buscar pedidos: " + err.Error())
	}
	defer rows.Close()

	// Mapeando os resultados
	purchasesMap := make(map[int]*ClientPurchases)

	for rows.Next() {
		var purchase Purchase
		var clientID int
		err := rows.Scan(&clientID, &purchase.PurchaseID, &purchase.PurchaseDate, &purchase.TotalAmount, &purchase.PaymentStatus, &purchase.PaymentMethod)
		if err != nil {
			return nil, err
		}

		// Verifica se o cliente já existe no mapa
		if clientPurchases, exists := purchasesMap[clientID]; exists {
			// Verifica se a compra já está na lista
			purchaseExists := false
			for _, p := range clientPurchases.Purchases {
				if p.PurchaseID == purchase.PurchaseID {
					purchaseExists = true
					break
				}
			}
			// Se a compra não existe, adicione
			if !purchaseExists {
				clientPurchases.Purchases = append(clientPurchases.Purchases, purchase)
			}
		} else {
			purchasesMap[clientID] = &ClientPurchases{
				ClientID:  clientID,
				Purchases: []Purchase{purchase},
			}
		}
	}

	var result []ClientPurchases
	for _, clientPurchases := range purchasesMap {
		result = append(result, *clientPurchases)
	}

	return result, nil
}
