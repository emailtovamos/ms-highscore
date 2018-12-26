// Command to run:  go run main.go  -address=localhost:50051

package main

import (
	// "io/ioutil"
	"flag"

	"github.com/rs/zerolog/log"
	pbhighscore "github.com/teach/ms-apis/ms-highscore/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func main() {

	var addressPtr = flag.String("address", "localhost:50051", "address to connect")

	flag.Parse()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure(), grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("addr", *addressPtr).Msg("failed to dial gRPC highscore service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("addr", *addressPtr).Msg("Failed to close the connection")
		}
	}()

	c := pbhighscore.NewGameClient(conn)

	if c == nil {
		log.Info().Msg("router client is nil")
	}

	r, err := c.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})

	if err != nil {
		log.Error().Err(err).Msg("Couldn't find path...")

	}

	if r != nil {
		log.Info().Interface("highscore: ", r.HighScore).Msg("")
	} else {
		log.Error().Msg("Couldn't get highscore.")
	}

}
