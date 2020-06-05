package handler

import (
	context "context"
	"fmt"
)

type Server struct {
}

func (*Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	fmt.Printf("ping successful %s", *in)
	return &PingMessage{Greeting: fmt.Sprintf("ping successful %s", in.Greeting)}, nil
}
