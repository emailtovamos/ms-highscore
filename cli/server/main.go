package main

import (
	"flag"
	// "google.golang.org/grpc"
	grpcSetup "github.com/emailtovamos/ms-highscore/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {

	var addressPtr = flag.String("address", "localhost:50051", "address where you can connect with ms-highscore service")

	flag.Parse()

	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	s := grpcSetup.NewServer(*addressPtr)

	// start gRPC server
	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("failed to start gRPC server")
	}

}
