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

const staticDir = "static"

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

	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	adminRealm := os.Getenv("ADMIN_REALM")

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

	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Simple handler for the root path
	mux.HandleFunc("/", cfg.handleHome)
	mux.HandleFunc("/home", cfg.handleHome)
	mux.HandleFunc("/article/{id...}", cfg.handleArticle)

	mux.HandleFunc("/admin", cfg.basicAuthMiddleware(cfg.handleAdmin, adminUsername, adminPassword, adminRealm))
	mux.HandleFunc("/new", cfg.basicAuthMiddleware(cfg.handleNew, adminUsername, adminPassword, adminRealm))
	mux.HandleFunc("/publish", cfg.basicAuthMiddleware(cfg.handlePublish, adminUsername, adminPassword, adminRealm))
	mux.HandleFunc("/edit/{id...}", cfg.basicAuthMiddleware(cfg.handleEdit, adminUsername, adminPassword, adminRealm))
	mux.HandleFunc("/update", cfg.basicAuthMiddleware(cfg.handleUpdate, adminUsername, adminPassword, adminRealm))
	mux.HandleFunc("/delete/{id...}", cfg.basicAuthMiddleware(cfg.handleDelete, adminUsername, adminPassword, adminRealm))

	http.ListenAndServe(":"+cfg.Port, mux)
}
