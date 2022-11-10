package model

import (
	"errors"
	"time"

	"github.com/edgedb/edgedb-go"
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
	Identifier  string    `json:"identifier" edgedb:"identifier"`
	CreatedBy   User      `json:"created_by" edgedb:"created_by"`
	OriginalURL string    `json:"original_url" edgedb:"original_url"`
	Visits      int64     `json:"visits" edgedb:"visits"`
	CreatedAt   time.Time `json:"created_at" edgedb:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" edgedb:"updated_at"`
}

type ShortenInput struct {
	RawURL     string
	Identifier Option[string]
	CreatedBy  User
}

type User struct {
	edgedb.Optional

	ID          edgedb.UUID `json:"id,omitempty" edgedb:"id"`
	IsActive    bool        `json:"is_verified,omitempty" edgedb:"is_active"`
	GitHubLogin string      `json:"gh_login" edgedb:"gh_login"`

	// TODO: should we store it in something like Vault?
	GitHubAccessKey string    `json:"gh_access_key,omitempty" edgedb:"gh_access_key"`
	CreatedAt       time.Time `json:"created_at,omitempty" edgedb:"created_at"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	User `json:"user_data"`
}
