package grpc

import (
	"context"
	pbhighscore "github.com/emailtovamos/ms-apis/ms-highscore/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 4445555.0

// NewServer creates a new instance of a gRPC server
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// GetHighScore returns the highscore from the HighScore variable
func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	// HighScore = 44444.0
	log.Info().Msg("GetHighScore in ms-highscore called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil

}

// SetHighScore sets the highscore to the HighScore variable
func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {

	HighScore = input.HighScore

	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil

}

// ListenAndServe starts the gRPC server on the given address
func (g *Grpc) ListenAndServe() error {
	// open tcp port to listen for incoming connections on
	// 1st Step
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open socket")
	}

	serverOpts := []grpc.ServerOption{}

	// create the server with the specified options
	g.srv = grpc.NewServer(serverOpts...)

	// 2nd Step
	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("addr", g.address).Msg("starting gRPC server")

	// start listening on the given address
	// 3rd Step
	if err := g.srv.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start gRPC server")
	}

	return nil
}
