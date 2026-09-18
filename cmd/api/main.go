package main

import (
	"log"
	"net/http"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/config"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/database"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/handler"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/repository"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/service"
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

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /post", postHandler.CreatePost)
	mux.HandleFunc("GET /posts", postHandler.GetPosts)
	mux.HandleFunc("UPDATE /post", postHandler.UpdatePost)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
