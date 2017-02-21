package handlers

import "database/sql"

// Env environment for handlers with database requirement
type Env struct {
	Conn *sql.DB
}
