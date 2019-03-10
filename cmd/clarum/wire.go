//go:generate wire
//+build wireinject

package main

import (
	"github.com/alee792/clarum/pkg/clarum"
	"github.com/alee792/clarum/pkg/storage/bolt"
	"github.com/google/wire"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type Config struct {
	Bolt struct {
		Path   string
		Bucket []byte
		Opts   *bbolt.Options
	}
}

var BoltSet = wire.NewSet(
	ProvideBolt,
	bolt.NewRepo,
	wire.Bind((*clarum.Repo)(nil), (*bolt.Repo)(nil)),
)

func ProvideBolt(cfg Config) (*bbolt.DB, error) {
	db, err := bbolt.Open(cfg.Bolt.Path, 0600, cfg.Bolt.Opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ProvideLogger(cfg Config) (*zap.SugaredLogger, error) {
	return zap.NewExample().Sugar(), nil
}

func InitClarum(cfg Config) (*clarum.Server, error) {
	wire.Build(
		ProvideLogger,
		BoltSet,
		clarum.NewServer,
		wire.FieldsOf(cfg, "Bolt.Bucket"),
	)
	return &clarum.Server{}, nil
}
