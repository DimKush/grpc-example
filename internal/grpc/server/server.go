package server

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"

	"github.com/DimKush/grpc-example/internal/config"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
)

type GRPCServer struct {
}

// NewGRPCServer creates a new GRPCServer
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (s *GRPCServer) Run(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()


	gatewayAddr := fmt.Sprintf("%s:%v", cfg.Rest.Host, cfg.Rest.Port)

	// metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)

	gatewayServer ;= create
}
