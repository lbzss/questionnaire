package data

import (
	"context"
	"questionnaire/ent"
	"questionnaire/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewQuestionnaireRepo, NewQuestionRepo)

// Data .
type Data struct {
	db   *ent.Client
	conf *conf.Data
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	entClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.NewHelper(logger).Fatalf("failed to open connection to db, %s", err)
		return nil, cleanup, err
	}
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.NewHelper(logger).Fatalf("failed to create schema, %s", err)
		return nil, cleanup, err
	}
	return &Data{
		db:   entClient,
		conf: c,
	}, cleanup, nil
}
