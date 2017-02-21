package models

import "database/sql"

// Client model
type Client struct {
	ID   int    `jsonapi:"primary,clients"`
	Name string `jsonapi:"attr,name"`
}

func getClient(rows *sql.Rows) (*Client, error) {
	item := new(Client)
	err := rows.Scan(&item.ID, &item.Name)
	return item, err
}

// AllClients returns all the clients
func AllClients(conn *sql.DB) ([]*Client, error) {
	rows, nok := conn.Query("SELECT * FROM clients")
	if nok != nil {
		return nil, nok
	}
	defer rows.Close()

	var results []*Client

	for rows.Next() {
		item, err := getClient(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}
	if nok = rows.Err(); nok != nil {
		return nil, nok
	}
	return results, nil
}

// FetchClient returns a single object found by ID
func FetchClient(conn *sql.DB, id int) (*Client, error) {
	rows, err := conn.Query("SELECT * FROM clients WHERE id=$1 LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	item, err := getClient(rows)
	if err != nil {
		return nil, err
	}

	return item, nil
}
