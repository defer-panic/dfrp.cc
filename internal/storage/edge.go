package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/edgedb/edgedb-go"
)

type edge struct {
	client *edgedb.Client
}

func NewEdgeDB(client *edgedb.Client) *edge {
	return &edge{client: client}
}

func (e *edge) Put(ctx context.Context, identifier, url string) (*model.Shortening, error) {
	const (
		op    = "storage.edge.Put"
		query = `insert Shortening {
			identifier := <str>$0,
			original_url := <str>$1,
			created_at := <datetime>$2
		};`
	)

	shortening := &model.Shortening{
		Identifier:  identifier,
		OriginalURL: url,
		Visits:      0,
		CreatedAt:   time.Now().UTC(),
	}

	var inserted struct{ id edgedb.UUID }

	if err := e.client.QuerySingle(
		ctx,
		query,
		&inserted,
		shortening.Identifier,
		shortening.OriginalURL,
		shortening.CreatedAt,
	); err != nil {
		var edbErr edgedb.Error
		if errors.As(err, &edbErr) {
			if edbErr.Category(edgedb.ConstraintViolationError) {
				return nil, model.ErrIdentifierExists
			}
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return shortening, nil
}

func (e *edge) Get(ctx context.Context, identifier string) (*model.Shortening, error) {
	const (
		op    = "storage.edge.Get"
		query = `select Shortening{identifier, original_url, created_at} filter .identifier = <str>$0`
	)

	var shortenings []model.Shortening

	if err := e.client.Query(ctx, query, &shortenings, identifier); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(shortenings) == 0 {
		return nil, model.ErrNotFound
	}

	return &shortenings[0], nil
}

func (e *edge) IncrementVisits(ctx context.Context, identifer string) error {
	const (
		op    = "storage.edge.IncrementVisits"
		query = `update Shortening
						 filter .identifier = <str>$0
						 set {
						 	visits := .visits + 1,
							updated_at := datetime_of_transaction()
						}`
	)

	if err := e.client.Execute(ctx, query, identifer); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
