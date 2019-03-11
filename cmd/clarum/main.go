package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/alee792/clarum/api"
	internal "github.com/alee792/clarum/internal/wire"
	"github.com/alee792/clarum/pkg/storage/bolt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	cfg := internal.Config{
		Bolt: &bolt.Config{
			Path:   "/tmp/test",
			Bucket: []byte("test"),
		},
	}
	c, err := internal.InitClarum(cfg)
	if err != nil {
		panic(err)
	}

	test := &pb.Instrument{
		Id:       "hello",
		IssuedAt: time.Now().UTC().Unix(),
	}
	ctx := context.Background()
	created, err := c.CreateInstrument(ctx, &pb.CreateInstrumentRequest{
		Instrument: test,
	})
	if err != nil {
		panic("Create failed: " + err.Error())
	}

	got, err := c.GetInstrument(ctx, &pb.GetInstrumentRequest{
		Id: test.Id,
	})
	if err != nil {
		panic("Get failed: " + err.Error())
	}

	if !cmp.Equal(created.IssuedAt, got.IssuedAt) || got.IssuedAt == 0 {
		panic(cmp.Diff(created, got))
	}

	fmt.Printf("successful retrieval: %+v\n", got)

	_, err = c.DeleteInstrument(ctx, &pb.DeleteInstrumentRequest{
		Id: test.Id,
	})
	if err != nil {
		panic(err)
	}

	empty, err := c.GetInstrument(ctx, &pb.GetInstrumentRequest{
		Id: test.Id,
	})
	if err != nil {
		panic(err)
	}

	if empty.Id != "" {
		fmt.Printf("there shouldn't be anything here... %+v\n", empty)
	}
}
