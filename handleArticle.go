package main

import (
	"html/template"
	"net/http"
)

func (cfg *Config) handleArticle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	articlePath := cfg.ArticleDirectory + "/article_" + id + ".json"
	article, err := LoadArticle(articlePath)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	data := struct {
		Article Article
	}{
		Article: article,
	}

	//fmt.Printf("article: %v %v %v", article.Title, article.Date_Published, article.Content)

	template, err := template.ParseFiles("./static/article.html")
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
