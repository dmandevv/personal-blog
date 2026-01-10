package main

import (
	"time"
)

type Article struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Date_Published time.Time `json:"date_published"`
}
