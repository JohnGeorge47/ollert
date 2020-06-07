package heartbeat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type serverdetails struct {
	Version    string  `json:"version"`
	Minversion string	`json:"minversion"`
	Buildtime  string	`json:"buildtime"`
	StartedAt  time.Time  `json:"started_at"`
	Uptime     string		`json:"uptime"`
}

var deets serverdetails

func init() {
	deets.StartedAt = time.Now()
}

func HeartBeat() http.Handler {
	uptime := time.Since(deets.StartedAt)
	formatedUptime := fmt.Sprintf("%02d:%02d:%02d", uptime.Hours(), uptime.Minutes(), uptime.Seconds())
	deets.Uptime = fmt.Sprintf("%d days %s", uptime/(time.Hour*24), formatedUptime)
	fn := func(w http.ResponseWriter, r *http.Request) {
		 jsonTosend,err:=json.Marshal(deets)
		 if err!=nil{
		 	w.WriteHeader(http.StatusInternalServerError)
		 	w.Write([]byte("Oops something went wrong"))
		 }
		 w.Write(jsonTosend)
	}
	return http.HandlerFunc(fn)
}
