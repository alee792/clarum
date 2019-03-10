// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package internal

import (
	"github.com/alee792/clarum/pkg/clarum"
	"github.com/alee792/clarum/pkg/storage/bolt"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InitClarum(config Config) (*clarum.Server, error) {
	boltConfig := config.Bolt
	repo, err := bolt.NewRepo(boltConfig)
	if err != nil {
		return nil, err
	}
	sugaredLogger, err := ProvideLogger(config)
	if err != nil {
		return nil, err
	}
	server := clarum.NewServer(repo, sugaredLogger)
	return server, nil
}

// wire.go:

type Config struct {
	Bolt *bolt.Config
}

var BoltSet = wire.NewSet(wire.FieldsOf(new(Config), "Bolt"), bolt.NewRepo, wire.Bind((*clarum.Repo)(nil), (*bolt.Repo)(nil)))

func ProvideLogger(cfg Config) (*zap.SugaredLogger, error) {
	return zap.NewExample().Sugar(), nil
}
