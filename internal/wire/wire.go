//go:generate wire
//+build wireinject

package internal

import (
	"github.com/alee792/clarum/pkg/clarum"
	"github.com/alee792/clarum/pkg/storage/bolt"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Config struct {
	Bolt *bolt.Config
}

var BoltSet = wire.NewSet(
	wire.FieldsOf(new(Config), "Bolt"),
	bolt.NewRepo,
	wire.Bind((*clarum.Repo)(nil), (*bolt.Repo)(nil)),
)

func ProvideLogger(cfg Config) (*zap.SugaredLogger, error) {
	return zap.NewExample().Sugar(), nil
}

func InitBolt(Config) (*bolt.Repo, error) {
	wire.Build(BoltSet)
	return nil, nil
}

func InitClarum(Config) (*clarum.Server, error) {
	wire.Build(
		clarum.NewServer,
		BoltSet,
		ProvideLogger,
	)
	return nil, nil
}
