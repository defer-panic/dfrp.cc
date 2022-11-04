package db

import (
	"context"

	"github.com/edgedb/edgedb-go"
)

type DB struct {
	client *edgedb.Client
}

func Connect(ctx context.Context, dsn string) (*DB, error) {
	client, err := edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{TLSSecurity: "insecure"})
	if err != nil {
		return nil, err
	}

	return &DB{client: client}, nil
}

func (d *DB) Client() *edgedb.Client {
	return d.client
}

func (d *DB) Close(ctx context.Context) error {
	var (
		closeCh = make(chan struct{})
		err     error
	)

	go func() {
		err = d.client.Close()
		closeCh <- struct{}{}
	}()

	select {
	case <-closeCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
