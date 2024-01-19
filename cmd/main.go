package main

import (
	"city2city/api"
	"city2city/api/handler"
	"city2city/config"
	"city2city/storage/postgres"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}

	defer store.CloseDB()

	handler := handler.New(store)

	api.New(handler)

	fmt.Println("Server is running on port 8080")
	if err = http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln("error while running server err:", err.Error())
	}
}