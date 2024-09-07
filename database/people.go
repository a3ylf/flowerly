package database

func (db *Database) GetClients() ([]Client, error) {
	rows, err := db.db.Query(`SELECT id, name, email, password, cpf, rua, num FROM client`)
		if err != nil {
		    return []Client{},err
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
	return clients,nil

}

