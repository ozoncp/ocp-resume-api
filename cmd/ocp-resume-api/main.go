package main

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-resume-api/internal/api"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"
)

const (
	grpcPort = ":82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpResumeApiServer(s, api.NewOcpResumeApi())

	if err := s.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	run()
}
