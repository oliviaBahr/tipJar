package core

import (
	"strings"

	"github.com/google/uuid"
)

type Tip struct {
	ID          string
	Title       string
	Description string
	Tags        []string
	Links       []string
}

func NewTip(title, description, tags, links string) *Tip {
	return &Tip{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Tags:        splitAndTrim(tags),
		Links:       splitAndTrim(links),
	}
}

func NewOldTip(id, title, description, tags, links string) *Tip {
	return &Tip{
		ID:          id,
		Title:       title,
		Description: description,
		Tags:        splitAndTrim(tags),
		Links:       splitAndTrim(links),
	}
}

func (t *Tip) Edit(title, description, tags, links string) {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if tags != "" {
		t.Tags = splitAndTrim(tags)
	}
	if links != "" {
		t.Links = splitAndTrim(links)
	}
}

func (t *Tip) FilterValue() string {
	return t.Title
}

func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
