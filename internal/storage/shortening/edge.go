package shortening

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/edgedb/edgedb-go"
)

const getQuery = `select Shortening{
			identifier,
			created_by: { id, gh_login, created_at },
			original_url,
			visits,
			created_at,
			updated_at
		} filter .identifier = <str>$0`

type edge struct {
	client *edgedb.Client
}

func NewEdgeDB(client *edgedb.Client) *edge {
	return &edge{client: client}
}

func (e *edge) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
	const (
		op    = "storage.edge.Put"
		query = `insert Shortening {
			identifier := <str>$0,
			original_url := <str>$1,
			created_by := (select User filter .gh_login = <str>$2)
		};`
	)

	var inserted struct{ id edgedb.UUID }

	if err := e.client.QuerySingle(
		ctx,
		query,
		&inserted,
		shortening.Identifier,
		shortening.OriginalURL,
		shortening.CreatedBy.GitHubLogin,
	); err != nil {
		var edbErr edgedb.Error
		if errors.As(err, &edbErr) {
			if edbErr.Category(edgedb.ConstraintViolationError) {
				return nil, model.ErrIdentifierExists
			}
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &shortening, nil
}

func (e *edge) Get(ctx context.Context, identifier string) (*model.Shortening, error) {
	const op = "storage.edge.Get"

	var shortenings []shortening

	if err := e.client.Query(ctx, getQuery, &shortenings, identifier); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(shortenings) == 0 {
		return nil, model.ErrNotFound
	}

	shorteningToReturn := toModelShortening(shortenings[0])
	return &shorteningToReturn, nil
}

func (e *edge) IncrementVisits(ctx context.Context, identifier string) error {
	const (
		op          = "storage.edge.IncrementVisits"
		updateQuery = `update Shortening
											 filter .identifier = <str>$0
											 set {
												visits := <int64>$1,
												updated_at := datetime_of_transaction()
											}`
	)

	if err := e.client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
		var s shortening
		if err := tx.QuerySingle(ctx, getQuery, &s, identifier); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if err := tx.Execute(ctx, updateQuery, identifier, s.Visits+1); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func fromModelShortening(m model.Shortening) shortening {
	return shortening{
		Identifier:  m.Identifier,
		CreatedBy:   m.CreatedBy,
		OriginalURL: m.OriginalURL,
		Visits:      m.Visits,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   edgedb.NewOptionalDateTime(m.UpdatedAt),
	}
}

func toModelShortening(s shortening) model.Shortening {
	updatedAt, _ := s.UpdatedAt.Get()

	return model.Shortening{
		Identifier:  s.Identifier,
		CreatedBy:   s.CreatedBy,
		OriginalURL: s.OriginalURL,
		Visits:      s.Visits,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   updatedAt,
	}
}

type shortening struct {
	Identifier  string                  `edgedb:"identifier"`
	CreatedBy   model.User              `edgedb:"created_by"`
	OriginalURL string                  `edgedb:"original_url"`
	Visits      int64                   `edgedb:"visits"`
	CreatedAt   time.Time               `edgedb:"created_at"`
	UpdatedAt   edgedb.OptionalDateTime `edgedb:"updated_at"`
}
