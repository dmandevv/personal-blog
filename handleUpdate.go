package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (cfg *Config) handleUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	title := r.FormValue("title")
	date := r.FormValue("date")
	content := r.FormValue("content")

	idInt := -1
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	dateParsed, _ := time.Parse("2006-01-02", date)

	var article = Article{
		ID:             idInt,
		Title:          title,
		Content:        content,
		Date_Published: dateParsed,
	}

	fileName := fmt.Sprintf("article_%d.json", idInt)

	err = SaveArticle(article, cfg.ArticleDirectory, fileName)
	if err != nil {
		http.Error(w, "Failed to save article", http.StatusInternalServerError)
		return
	}
	cfg.SaveConfig()

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
