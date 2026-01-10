package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ArticleDirectory string
	Port             string
	NextArticleID    int `json:"next_article_id"`
}

func main() {
	godotenv.Load()
	articleDir, exists := os.LookupEnv("ARTICLE_DIRECTORY")
	if !exists || articleDir == "" {
		fmt.Println("Missing ARTICLE_DIRECTORY")
	}
	port, exists := os.LookupEnv("PORT")
	if !exists || port == "" {
		fmt.Println("Missing PORT")
	}

	cfg, err := LoadConfig()
	if err != nil {
		fmt.Printf("%v - defaulting to .env values\n", err)
		cfg = &Config{
			ArticleDirectory: articleDir,
			Port:             port,
			NextArticleID:    1,
		}
	} else {
		// Override loaded config with environment variables
		cfg.ArticleDirectory = articleDir
		cfg.Port = port
	}

	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Simple handler for the root path
	http.HandleFunc("/", cfg.handleHome)
	http.HandleFunc("/home", cfg.handleHome)
	http.HandleFunc("/admin", cfg.handleAdmin)
	http.HandleFunc("/new", cfg.handleNew)
	http.HandleFunc("/publish", cfg.handlePublish)
	http.HandleFunc("/article/{id}/", cfg.handleArticle)

	http.ListenAndServe(":"+cfg.Port, nil)
}
