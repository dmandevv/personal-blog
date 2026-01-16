package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func (cfg *Config) handleEdit(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	articlePath := filepath.Join(cfg.ArticleDirectory, "article_"+id+".json")
	article, err := LoadArticle(articlePath)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	data := struct {
		Article    Article
		StringDate string
	}{
		Article:    article,
		StringDate: article.Date_Published.Format("2006-01-02"),
	}

	tmpl, err := template.ParseFiles("./static/edit.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute template:%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(buf.Bytes())
}
