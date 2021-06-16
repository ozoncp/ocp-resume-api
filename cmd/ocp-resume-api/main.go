package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-resume-api/internal/api"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"
)

func run(grpcPort *string,
	sqlPort *string,
	sqlUser *string,
	sqlPassword *string,
	sqlDbname *string) error {

	ctx := context.Background()
	listen, err := net.Listen("tcp", ":"+*grpcPort)
	if err != nil {
		log.Err(err).Msg("failed to create grpc listener")
		return err
	}

	psqlInfo := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", *sqlPort, *sqlUser, *sqlPassword, *sqlDbname)
	db, err := sqlx.Open("pgx", psqlInfo)

	if err != nil {
		log.Error().Err(err).Msg("failed to open database")
		return err
	}
	defer db.Close()
	err = db.PingContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to ping database")
		return err
	}
	repo := repo.NewRepo(*db)

	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpResumeApiServer(s, api.NewOcpResumeApi(repo))

	if err := s.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	grpcPort := flag.String("grpc_port", "82", "Define port number for service.")
	sqlPort := flag.String("sql_port", "5432", "Define SQL port number.")
	sqlUser := flag.String("sql_user", "postgres", "Define SQL username")
	sqlPassword := flag.String("sql_pass", "123", "Define SQL password")
	sqlDbname := flag.String("sql_dbname", "postgres", "Define SQL DB name")
	flag.Parse()
	if err := run(grpcPort, sqlPort, sqlUser, sqlPassword, sqlDbname); err == nil {
		log.Err(err).Msg("")
	}
}
