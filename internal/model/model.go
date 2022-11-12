package model

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	. "github.com/samber/mo"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
	ErrInvalidURL       = errors.New("invalid url")
	ErrUserIsNotMember  = errors.New("user is not member of the organization")
	ErrInvalidToken     = errors.New("invalid token")
)

type Shortening struct {
	Identifier  string    `json:"identifier"`
	CreatedBy   string    `json:"created_by"`
	OriginalURL string    `json:"original_url"`
	Visits      int64     `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ShortenInput struct {
	RawURL     string
	Identifier Option[string]
	CreatedBy  string
}

type User struct {
	IsActive    bool        `json:"is_verified,omitempty"`
	GitHubLogin string      `json:"gh_login"`

	// TODO: should we store it in something like Vault?
	GitHubAccessKey string    `json:"gh_access_key,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	User `json:"user_data"`
}
