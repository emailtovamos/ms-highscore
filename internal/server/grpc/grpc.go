package grpc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	pbhighscore "github.com/teach/ms-apis/ms-highscore/v1"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

// var (
// 	HighScore float64
// )
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
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open socket")
	}

	// tracers will default to a NOOP tracer if nothing was configured
	// streamTracingInterceptor := grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer()))
	// unaryTracingInterceptor := grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer()))

	// add prometheus middleware
	// grpc_prometheus.EnableHandlingTimeHistogram(grpc_prometheus.WithHistogramBuckets(
	// 	metrics_helpers.ExpBuckets(
	// 		float64(10*time.Microsecond)/float64(time.Second),
	// 		1, 20,
	// 	),
	// ))

	// serverOpts := []grpc.ServerOption{
	// 	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	// 		streamTracingInterceptor,
	// 		grpc_prometheus.StreamServerInterceptor,
	// 	)),
	// 	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	// 		unaryTracingInterceptor,
	// 		grpc_prometheus.UnaryServerInterceptor,
	// 	))}
	serverOpts := []grpc.ServerOption{}

	// create the server with the specified options
	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)

	// initialize health server
	// hs := health.NewServer()
	// grpc_health_v1.RegisterHealthServer(g.srv, hs)

	log.Info().Str("addr", g.address).Msg("starting gRPC server")

	// shutdown the server gracefully when leaving this function
	// defer func() {
	// 	g.GracefulStop()
	// }()

	// start listening on the given address
	if err := g.srv.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start gRPC server")
	}

	return nil
}
