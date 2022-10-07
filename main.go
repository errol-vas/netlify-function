package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type RemoteAddr struct {
	LocalIP  string `json:"local_ip"`
	PublicIP string `json:"public_ip"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	resp, err := http.Get("https://icanhazip.com/")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	defer resp.Body.Close()
	publicIP, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	ip := RemoteAddr{
		LocalIP:  r.RemoteAddr,
		PublicIP: string(publicIP),
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
