package sql

import (
	"database/sql"
	"github.com/JohnGeorge47/ollert/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

var Connmanager Isql

func init() {
	err := godotenv.Load()
	c := config.Get()
	db, err := sql.Open("mysql", c.Connectionstring())
	if err != nil {
		log.Fatal("Error in connecting to mysql", err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(3)
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to mysql")
	}
	Connmanager = Client{DB: db}
}
