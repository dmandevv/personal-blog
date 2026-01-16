package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (a *Article) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}

func ArticleFromJSON(data []byte) (Article, error) {
	var a Article
	err := json.Unmarshal(data, &a)
	if err != nil {
		return a, err
	}
	return a, nil
}

func SaveArticle(article Article, directoryPath, fileName string) error {
	// Ensure the directory exists
	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := article.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert article to JSON: %w", err)
	}

	filePath := filepath.Join(directoryPath, fileName)
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

func LoadArticle(filePath string) (Article, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Article{}, fmt.Errorf("failed to read JSON file: %w", err)
	}

	article, err := ArticleFromJSON(data)
	if err != nil {
		return Article{}, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return article, nil
}

func (cfg *Config) LoadAllArticles() []Article {
	var articles []Article

	files, err := os.ReadDir(cfg.ArticleDirectory)
	if err != nil {
		fmt.Printf("failed to read article directory: %v\n", err)
		return articles
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(cfg.ArticleDirectory, file.Name())
		article, err := LoadArticle(filePath)
		if err != nil {
			fmt.Printf("failed to load article from %s: %v\n", filePath, err)
			continue
		}

		articles = append(articles, article)
	}

	return articles
}

func (cfg *Config) SaveConfig() error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile("./config.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %w", err)
	}

	return &cfg, nil
}
