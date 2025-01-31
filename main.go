package main

import (
	controller "Cash-Flow-Apps/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// for load godotenv
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/cash-flow", controller.GetCashFlow).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)
	log.Println("Starting on Port 8080")

	err := http.ListenAndServe(":8080", handler)
	log.Fatal(err)
}
