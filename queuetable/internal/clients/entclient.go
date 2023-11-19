package clients

import (
	"context"
	"queuetable/ent"
	"queuetable/internal/env"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func EntClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(env.EntDriver, env.EntDatasourceName)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening connection to database")
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed creating schema resources")
	}

	return client.Debug(), err
}
