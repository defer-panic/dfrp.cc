package shorten

import (
	"context"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/google/uuid"
)

type Storage interface {
	Put(ctx context.Context, identifier, url string) (*model.Shortening, error)
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(ctx context.Context, input model.ShortenInput) (*model.Shortening, error) {
	var (
		id         = uuid.New().ID()
		identifier = input.Identifier.OrElse(Shorten(id, input.RawURL))
	)

	shortening, err := s.storage.Put(ctx, identifier, input.RawURL)
	if err != nil {
		return nil, err
	}

	return shortening, nil
}

func (s *Service) GetRedirectURL(ctx context.Context, identifier string) (string, error) {
	shortening, err := s.storage.Get(ctx, identifier)
	if err != nil {
		return "", err
	}

	return shortening.OriginalURL, nil
}
