package server

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net/http"
	"runtime"
)

func createGatewayServer(grpcAddr, gatewayAddr string) *http.Server {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
		),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial server")
	}

	mux := runtime.NewServeMux()
	//if err := pb.RegisterOmpTemplateApiServiceHandler(context.Background(), mux, conn); err != nil {
	//	log.Fatal().Err(err).Msg("Failed registration handler")
	//}

	gatewayServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: tracingWrapper(mux),
	}

	return gatewayServer
}
