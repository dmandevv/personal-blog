package main

import (
	"net/http"
)

func (cfg *Config) handleAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/admin.html")
}
