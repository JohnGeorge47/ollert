package sql


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/JohnGeorge47/ollert/internal/config"
	"log"
)

var Connmanager Isql

func init(){
	c:=config.Get()
	db,err:=sql.Open("mysql",c.Connectionstring())
	if err!=nil{
		log.Fatal("Error in connecting to mysql",err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(3)
	Connmanager=Client{DB:db}
}