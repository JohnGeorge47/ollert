package sql

import (
	"database/sql"
)

type Client struct {
	DB *sql.DB
}
