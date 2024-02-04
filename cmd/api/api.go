package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	var r *chi.Mux = chi.NewRouter()

	fmt.Println("Starting the api")

	err := http.ListenAndServe("localhost:5000", r)
	if err != nil {
		log.Fatal(err)
	}

}