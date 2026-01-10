package main

import (
	"fmt"
	"net/http"
	"time"
)

func (cfg *Config) handlePublish(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	date := r.FormValue("date")
	content := r.FormValue("content")

	dateParsed, _ := time.Parse("2006-01-02", date)

	var newArticle = Article{
		Title:          title,
		Content:        content,
		Date_Published: dateParsed,
	}

	fileName := fmt.Sprintf("article_%d.json", cfg.NextArticleID)

	err := SaveArticle(newArticle, cfg.ArticleDirectory, fileName)
	if err != nil {
		http.Error(w, "Failed to save article", http.StatusInternalServerError)
		return
	}
	cfg.NextArticleID++
	cfg.SaveConfig()

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
