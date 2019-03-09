package bolt

import (
	"context"

	pb "github.com/alee792/clarum/api"
	"github.com/golang/protobuf/proto"
	bolt "go.etcd.io/bbolt"
)

// Config for Repo.
type Config struct {
	BoltFilePath string
}

// Repo stores Clarum data using BoltDB.
type Repo struct {
	db     bolt.DB
	bucket []byte
}

// CreateInstrument to begin tracking lineage.
func (r *Repo) CreateInstrument(ctx context.Context, in *pb.CreateInstrumentRequest) (*pb.Instrument, error) {
	return r.update(ctx, in.Instrument.Id, in.Instrument)
}

// GetInstrument returns an instrument.
func (r *Repo) GetInstrument(ctx context.Context, in *pb.GetInstrumentRequest) (*pb.Instrument, error) {
	var inst *pb.Instrument
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		raw := b.Get([]byte(in.Id))
		return proto.Unmarshal(raw, inst)
	})
	if err != nil {
		return nil, err
	}
	return inst, nil
}

// ListInstruments returns a list of instruments.
func (r *Repo) ListInstruments(ctx context.Context, in *pb.ListInstrumentsRequest) (*pb.ListInstrumentsResponse, error) {
	var insts []*pb.Instrument
	err := r.db.View(func(tx *bolt.Tx) error {
		for _, id := range in.Ids {
			b := tx.Bucket(r.bucket)
			raw := b.Get([]byte(id))
			var inst *pb.Instrument
			if err := proto.Unmarshal(raw, inst); err != nil {
				return err
			}
			insts = append(insts, inst)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteInstrument removes an instrument from a server.
func (r *Repo) DeleteInstrument(ctx context.Context, in *pb.DeleteInstrumentRequest) (*pb.Empty, error) {
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		return b.Delete([]byte(in.Id))
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

// UpdateInstrument updates an instrument with the requests validated fields.
func (r *Repo) UpdateInstrument(ctx context.Context, in *pb.UpdateInstrumentRequest) (*pb.Instrument, error) {
	return r.update(ctx, in.Id, in.Instrument)
}

func (r *Repo) update(ctx context.Context, id string, inst *pb.Instrument) (*pb.Instrument, error) {
	raw, err := proto.Marshal(inst)
	if err != nil {
		return nil, err
	}
	err = r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		return b.Put([]byte(id), raw)
	})
	if err != nil {
		return nil, err
	}
	return inst, nil
}
