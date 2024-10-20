package database

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
