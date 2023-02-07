package receiver

import (
	"context"

	database "github.com/raythx98/go-database/sqlc/output"
)

type Database struct {
	Queries *database.Queries
}

var DbInstance Database

func (db *Database) AddLink(ctx context.Context, addLinkParams database.AddLinkParams) error {
	return db.Queries.AddLink(ctx, addLinkParams)
}

func (db *Database) GetFullLink(ctx context.Context, shortenedLink string) (string, error) {
	return db.Queries.GetFullLink(ctx, shortenedLink)
}
