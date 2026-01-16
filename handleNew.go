package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func (cfg *Config) handleNew(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Date string
	}{
		Date: time.Now().Format("2006-01-02"),
	}

	tmpl, err := template.ParseFiles("./static/new.html")
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
