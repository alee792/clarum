//+build wireinject

package main

import (
	"github.com/alee792/clarum/pkg/clarum"
	"github.com/alee792/clarum/pkg/storage/bolt"
	"github.com/google/wire"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

// Config for Resolver.
type Config struct {
	Bolt struct{ 
		Path string
		Opts bbolt.Options
	}
}  
 
// Resovler of dependencies.
type Resolver struct { 
	bolt *bbolt.DB
}   

func provideBolt() (*bbolt.DB, err)  {
	if r.bolt == nil {
		db ,err := bbolt.Open(r.Config.Bolt.Path, 0600, r.Config.Bolt.Opts)
		if err != nil {
			return nil, err
		} 
		r.bolt = db  
	} 
	return r.bolt, nil
}

func initClarum() *clarum.Server {
	wire.Build(
		zap.NewDevelopment,
		bolt.NewRepo,
		wire.Bind((*clarum.Repo)(nil), (*bolt.Repo)(nil)),
		clarum.NewServer,
	)
	return &clarum.Server{}
}
