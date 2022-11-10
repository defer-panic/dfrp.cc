package user

import (
	"context"
	"errors"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/edgedb/edgedb-go"
)

type edge struct {
	client *edgedb.Client
}

func NewEdgeDB(edgeClient *edgedb.Client) *edge {
	return &edge{client: edgeClient}
}

func (e *edge) CreateOrUpdate(ctx context.Context, user model.User) (*model.User, error) {
	const query = `insert User {
		gh_login := <str>$0,
		gh_access_key := <str>$1,
	}
	unless conflict on .gh_login
	else (
		update User
		set { gh_access_key := <str>$1 }
	);`
	var inserted model.User

	if err := e.client.QuerySingle(ctx, query, &inserted, user.GitHubLogin, user.GitHubAccessKey); err != nil {
		return nil, err
	}
	user.ID = inserted.ID

	return &user, nil
}

func (e *edge) Update(ctx context.Context, user model.User) error {
	const query = `update User 
								 filter .gh_login = <str>$0
								 set {
									 gh_access_key := <str>$1
								 }`

	if err := e.client.Execute(ctx, query, user.GitHubLogin, user.GitHubAccessKey); err != nil {
		return err
	}

	return nil
}

func (e *edge) GetByGithubLogin(ctx context.Context, login string) (*model.User, error) {
	const query = `select User { is_active, gh_login, gh_access_key, created_at} filter .gh_login = <str>$0;`

	var user model.User
	if err := e.client.QuerySingle(ctx, query, &user, login); err != nil {
		var edbErr edgedb.Error
		if errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (e *edge) Deactivate(ctx context.Context, login string) error {
	const query = `update User set { is_active := false } where .gh_login = <str>$0;`

	if err := e.client.Execute(ctx, query, login); err != nil {
		var edbErr edgedb.Error
		if errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError) {
			return model.ErrNotFound
		}
		return err
	}

	return nil
}
