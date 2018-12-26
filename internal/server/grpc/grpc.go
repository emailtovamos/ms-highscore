package grpc

import (
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

// NewServer creates a new instance of a gRPC server
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}
