package main

import (
	"flag"
	"fmt"
	"github.com/JohnGeorge47/ollert/internal/config"
	"github.com/JohnGeorge47/ollert/internal/core/accountCreation"
	"github.com/JohnGeorge47/ollert/internal/core/heartbeat"
	_ "github.com/JohnGeorge47/ollert/pkg/sql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

var port = flag.String("port", "3003", "You know the port or something")

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error while loading env variables", err)
	}
	c := config.Get()
	fmt.Println(c.MySqlConfig.DbPwd)
	r := mux.NewRouter()
	r.Handle("/", heartbeat.HeartBeat())
	r.Handle("/createAccount", accountCreation.Create())
	err = http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
	if err != nil {
		log.Fatal(err)
	}
}
