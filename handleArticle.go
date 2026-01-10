package main

import (
	"net/http"
)

func (cfg *Config) handleArticle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/article.html")
}
