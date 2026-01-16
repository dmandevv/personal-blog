package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func (cfg *Config) basicAuthMiddleware(next http.HandlerFunc, expectedUsername, expectedPassword, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(expectedUsername))
			expectedPasswordHash := sha256.Sum256([]byte(expectedPassword))

			validUsername := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			validPassword := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if validUsername && validPassword {
				next(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
