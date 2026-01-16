package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func (cfg *Config) handleHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Articles []Article
	}{
		Articles: cfg.LoadAllArticles(),
	}

	tmpl, err := template.ParseFiles("./static/index.html")
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
