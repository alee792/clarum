// +build wireinject

package main

import (
	"github.com/alee792/clarum/pkg/clarum"
	"github.com/alee792/clarum/pkg/storage/bolt"
	"github.com/google/wire"
)

func initClarum(db *bolt.Repo) *clarum.Server {
	wire.Build(bolt.NewRepo)
	return &clarum.Server{}
}
