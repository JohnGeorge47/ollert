package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/JohnGeorge47/ollert/internal/core/heartbeat"
	"github.com/JohnGeorge47/ollert/internal/config"
	"github.com/joho/godotenv"
)

var port = flag.String("port", "3003", "You know the port or something")

func main() {
	flag.Parse()
	err:=godotenv.Load()
	if err!=nil{
		fmt.Println("Error while loading env variables",err)
	}
	c:=config.Get()
	fmt.Println(c.MySqlConfig.DbPwd)
	r := mux.NewRouter()
	r.Handle("/",heartbeat.HeartBeat())
	err = http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
	if err != nil {
		log.Fatal(err)
	}
}
