package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diommyliel/GoAPI/internal/handlers"
	"github.com/go-chi/chi"
)

func main() {

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting the api")

	err := http.ListenAndServe("localhost:5000", r)
	if err != nil {
		log.Fatal(err)
	}

}
