package main

import (
	"html/template"
	"net/http"
)

func (cfg *Config) handleHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Articles []Article
	}{
		Articles: cfg.LoadAllArticles(),
	}

	template, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}
