package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	// set headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	godotenv.Load()
	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
