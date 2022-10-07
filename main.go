package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RemoteAddr struct {
	Ip string `json:"ip"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ip := RemoteAddr{
		Ip: r.RemoteAddr,
	}

	js, err := json.Marshal(ip)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
