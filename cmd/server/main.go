package server

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var port = flag.String("port", "8000", "You know the port or something")

func main() {
	flag.Parse()
	r := mux.NewRouter()
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal(err)
	}
}
