package main

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

const (
	ADDR = ":8000"
)

type product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", productsHandler)
	http.Handle("/", r)

	log.Println("API Server Started on", ADDR)
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		log.Fatalln(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)

	body := "Hello from home v0.9"
	log.Println("Home 0.9")

	json.NewEncoder(w).Encode(body)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)

	log.Println("Products 0.2")

	products := []product{
		{1, "XBOX", "200.02"},
		{2, "Playstation", "300.03"},
		{3, "Wii", "100.01"},
		{4, "Atari", "5.00"},
	}

	json.NewEncoder(w).Encode(products)
}

