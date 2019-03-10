package main

import (
	"fmt"

	wire "github.com/alee792/clarum/internal/wire"
	"github.com/alee792/clarum/pkg/storage/bolt"
)

func main() {
	cfg := wire.Config{
		Bolt: &bolt.Config{
			Path:   "/tmp/test",
			Bucket: []byte("test"),
		},
	}
	c, err := wire.InitClarum(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
