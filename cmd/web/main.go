package main

import (
	"log"
	"os"

	"github.com/Oloruntobi1/hubucweb/internal/handler"
	"github.com/Oloruntobi1/hubucweb/internal/repository"
)

func init() {
	// wire up config using the util library
}

func main() {

	// get current storage service from env variable

	storageService := os.Getenv("CURRENT_STORAGE")

	db := repository.GetStorage(storageService)

	server, err := handler.NewServer(db)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// start the server
	err = server.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
