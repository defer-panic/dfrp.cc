package storage

import (
	"context"
	"sync"
	"time"

	"github.com/defer-panic/url-shortener-api/internal/model"
)

type inMemory struct {
	m sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (s *inMemory) Put(_ context.Context, identifier string, url string) (*model.Shortening, error) {
	shortening := model.Shortening{
		Identifier:  identifier,
		OriginalURL: url,
		CreatedAt:   time.Now().UTC(),
	}

	s.m.Store(identifier, shortening)

	return &shortening, nil
}

func (s *inMemory) Lookup(_ context.Context, url string) (bool, error) {
	var found bool

	// ok for tests, but not for production
	s.m.Range(func(key, value any) bool {
		shortening := value.(model.Shortening)

		if shortening.OriginalURL == url {
			found = true
			return false
		}

		return true
	})

	return found, nil
}

func (s *inMemory) Get(_ context.Context, identifier string) (*model.Shortening, error) {
	v, ok := s.m.Load(identifier)
	if !ok {
		return nil, model.ErrNotFound
	}

	shortening := v.(model.Shortening)

	return &shortening, nil
}
