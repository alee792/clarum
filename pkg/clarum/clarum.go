package clarum

import (
	"context"

	pb "github.com/alee792/clarum/api"
	"go.uber.org/zap"
)

// Repo provides access to instruments.
// It uses the same protobuf defined interface as Clarum itself.
type Repo interface {
	pb.InstrumentServiceServer
}

// Server is an instance of Clarum that performs instrument management.
type Server struct {
	Repo
	Logger *zap.SugaredLogger
}

// NewServer creates a Clarum service from its dependencies.
func NewServer(repo Repo, logger *zap.SugaredLogger) *Server {
	s := &Server{
		Repo:   repo,
		Logger: logger,
	}
	return s
}

// CreateInstrument to begin tracking lineage.
func (s *Server) CreateInstrument(ctx context.Context, in *pb.CreateInstrumentRequest) (*pb.Instrument, error) {
	return s.Repo.CreateInstrument(ctx, in)
}

// GetInstrument returns an instrument.
func (s *Server) GetInstrument(ctx context.Context, in *pb.GetInstrumentRequest) (*pb.Instrument, error) {
	return s.Repo.GetInstrument(ctx, in)
}

// ListInstruments returns a list of instruments.
func (s *Server) ListInstruments(ctx context.Context, in *pb.ListInstrumentsRequest) (*pb.ListInstrumentsResponse, error) {
	return s.Repo.ListInstruments(ctx, in)
}

// DeleteInstrument removes an instrument from a server.
func (s *Server) DeleteInstrument(ctx context.Context, in *pb.DeleteInstrumentRequest) (*pb.Empty, error) {
	return s.Repo.DeleteInstrument(ctx, in)
}

// UpdateInstrument updates an instrument with the requests validated fields.
func (s *Server) UpdateInstrument(ctx context.Context, in *pb.UpdateInstrumentRequest) (*pb.Instrument, error) {
	return s.Repo.UpdateInstrument(ctx, in)
}
