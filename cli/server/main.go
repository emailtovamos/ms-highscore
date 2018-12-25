package main

import (
	"flag"
	// "google.golang.org/grpc"
)

func main() {

	var addressPtr = flag.String("address", "localhost:50051", "address where you can connect with ms-highscore service")

	flag.Parse()

	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

}
