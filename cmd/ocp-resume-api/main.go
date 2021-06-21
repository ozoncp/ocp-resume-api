package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/ozoncp/ocp-resume-api/internal/metrics"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/ozoncp/ocp-resume-api/internal/api"
	"github.com/ozoncp/ocp-resume-api/internal/configs"
	"github.com/ozoncp/ocp-resume-api/internal/producer"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"
)

func runMetrics(metricsConfig configs.Metrics) {
	metrics.Register()
	http.Handle(metricsConfig.Pattern, promhttp.Handler())
	go func() {
		err := http.ListenAndServe(metricsConfig.Address, nil)
		if err != nil {
			panic(err)
		}
	}()
}

func openDatabase(dbConfigs configs.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfigs.Port, dbConfigs.User, dbConfigs.Password, dbConfigs.Database)

	return sqlx.Open("pgx", dsn)
}

func runGrpc(
	grpcConfigs configs.Grpc,
	dbConfigs configs.Database,
	kafkaConfigs configs.Kafka,
) error {

	listener, err := net.Listen("tcp", grpcConfigs.Address)
	if err != nil {
		return fmt.Errorf("failed to start listening: %v", err)
	}

	db, err := openDatabase(dbConfigs)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Warn().Err(err).Msg("Database is inaccessable")
	}

	prod := producer.New(kafkaConfigs)
	prod.Init()
	defer prod.Close()

	server := grpc.NewServer()
	reflection.Register(server)
	api := api.NewOcpResumeApi(repo.NewRepo(db, dbConfigs.BatchSize), prod)
	desc.RegisterOcpResumeApiServer(server, api)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve server: %v", err)
	}

	return nil
}

func main() {
	cfg, err := configs.Read("config.yml")
	if err != nil {
		log.Fatal().Msgf("failed to open configuration file: %v", err)
		return
	}

	runMetrics(cfg.Metrics)
	if err := runGrpc(cfg.Grpc, cfg.Database, cfg.Kafka); err != nil {
		log.Fatal().Msgf("failed to start GRPC server: %v", err)
	}
}
