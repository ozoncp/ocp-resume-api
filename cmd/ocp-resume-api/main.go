package main

import (
	"flag"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-resume-api/internal/api"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"
)

func run(portNum *string) error {
	listen, err := net.Listen("tcp", ":"+*portNum)
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
	portNum := flag.String("port", "82", "Define port number for service.")
	flag.Parse()
	if err := run(portNum); err == nil {
		log.Fatal()
	}
}
