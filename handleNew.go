package main

import (
	"html/template"
	"net/http"
	"time"
)

func (cfg *Config) handleNew(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/new.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Date string
	}{
		Date: time.Now().Format("2006-01-02"),
	}
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}
