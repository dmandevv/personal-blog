package main

import (
	"net/http"
)

func (cfg *Config) handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
