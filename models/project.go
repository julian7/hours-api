package models

import (
	"database/sql"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Project model
type Project struct {
	ID       int       `jsonapi:"primary,projects"`
	ClientID int       `jsonapi:"attr,client_id"`
	Name     string    `jsonapi:"attr,name"`
	Start    time.Time `jsonapi:"attr,start"`
	Finish   time.Time `jsonapi:"attr,finish"`
}

func getProject(rows *sql.Rows) (*Project, error) {
	item := new(Project)
	err := rows.Scan(&item.ID, &item.ClientID, &item.Name, &item.Start, &item.Finish)
	return item, err
}

// AllProjects returns all the clients
func AllProjects(conn *sql.DB, filters url.Values) ([]*Project, error) {
	query := "SELECT * FROM projects"
	var where []string
	var params []interface{}
	i := 0
	filterRegex := regexp.MustCompile("^filter\\[([a-z_]+)\\]$")
	for setting := range filters {
		match := filterRegex.FindStringSubmatch(setting)
		if match != nil {
			i = i + 1
			where = append(where, fmt.Sprintf("%s=$%d", match[1], i))
			params = append(params, filters[setting][0])
		}
	}
	if len(where) > 0 {
		query = query + " WHERE " + strings.Join(where, " AND ")
	}
	rows, nok := conn.Query(query, params...)
	if nok != nil {
		return nil, nok
	}
	defer rows.Close()

	var results []*Project

	for rows.Next() {
		item, err := getProject(rows)
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

// FetchProject returns a single object found by ID
func FetchProject(conn *sql.DB, id int) (*Project, error) {
	rows, err := conn.Query("SELECT * FROM projects WHERE id=$1 LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	item, err := getProject(rows)
	if err != nil {
		return nil, err
	}

	return item, nil
}
