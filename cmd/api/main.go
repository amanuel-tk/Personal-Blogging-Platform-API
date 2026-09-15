package main

import (
	"fmt"
	"log"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/config"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/database"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDB(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to PostgresSQL")
}
