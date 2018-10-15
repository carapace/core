package postgres

import (
	"database/sql"
)

type Engine struct {
	db *sql.DB
}
