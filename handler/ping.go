package handler

import (
	context "context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "requests_total",
		Help: "total number of requests",
	})
)

type Server struct {
}

func (*Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	fmt.Printf("ping successful %s", *in)
	requestCounter.Inc()
	return &PingMessage{Greeting: fmt.Sprintf("ping successful %s", in.Greeting)}, nil
}
