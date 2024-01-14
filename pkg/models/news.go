// pkg/models/news.go
package models

import _ "errors"

// News represents a news article.
type News struct {
	ID      int
	Title   string
	Content string
}

// NewsRepository is an interface for interacting with news data.
type NewsRepository interface {
	GetAll() ([]News, error)
	AddArticle(title, content string) error
}

// MemoryNewsRepository is an in-memory implementation of NewsRepository.
type MemoryNewsRepository struct {
	news []News
}

func (r *MemoryNewsRepository) GetAll() ([]News, error) {
	return r.news, nil
}

func (r *MemoryNewsRepository) AddArticle(title, content string) error {
	// Implement the logic to add articles to the in-memory repository
	// This is optional if you still want to use in-memory storage for some parts
	return nil
}
