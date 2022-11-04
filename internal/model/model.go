package model

import (
	"errors"
	"time"

	. "github.com/samber/mo"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
	ErrInvalidURL       = errors.New("invalid url")
)

type Shortening struct {
	Identifier  string    `json:"identifier" edgedb:"identifier"`
	OriginalURL string    `json:"original_url" edgedb:"original_url"`
	CreatedAt   time.Time `json:"created_at" edgedb:"created_at"`
}

type ShortenInput struct {
	RawURL     string
	Identifier Option[string]
}
