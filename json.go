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

func ArticleFromJSON(data []byte) (*Article, error) {
	var a Article
	err := json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func SaveArticle(article Article, directoryPath, fileName string) error {

	// Ensure the directory exists
	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := filepath.Join(directoryPath, fileName)

	jsonData, err := article.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert article to JSON: %w", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

func LoadArticle(filePath string) (*Article, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	article, err := ArticleFromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return article, nil
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
