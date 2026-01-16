package main

import (
	"net/http"
	"path/filepath"
)

func (cfg *Config) handleDelete(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	articlePath := filepath.Join(cfg.ArticleDirectory, "article_"+id+".json")

	err := DeleteArticle(articlePath)

	//Deletion failed
	if err != nil {
		http.Error(w, "Failed to delete article", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
