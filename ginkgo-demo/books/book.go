package books

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID            string
	Title         string
	Author        string
	PublishedDate *time.Time
}

func NewBook(title, author string, date *time.Time) *Book {
	return &Book{
		ID:            uuid.New().String(),
		Title:         title,
		Author:        author,
		PublishedDate: date,
	}
}

func (b *Book) Description() string {
	if len(b.Title) == 0 {
		return "No Title"
	}
	display := b.Title
	if len(b.Author) > 0 {
		display = fmt.Sprintf("%s, %s", display, b.Author)
	}
	if b.PublishedDate != nil {
		display = fmt.Sprintf("%s %s", display, b.PublishedDate)
	}
	return display
}

func (b *Book) IsNewBook() bool {
	if b.PublishedDate == nil {
		return false
	}
	newPublishTime, _ := time.Parse("2006-10-02", "2023-01-01")
	return b.PublishedDate.Unix() >= newPublishTime.Unix()
}
