package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/MaksimYudenko/training/pkg/protocol/grpc"
	"github.com/MaksimYudenko/training/pkg/protocol/rest"
	"github.com/MaksimYudenko/training/pkg/service/v1"
	_ "github.com/lib/pq"
)

// Configuration for Server
type Config struct {
	// gRPC server starts parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string
	// DB parameters section
	dbHost     string
	dbPort     int
	dbName     string
	dbUser     string
	dbPassword string
	dbSchema   string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()
	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.dbHost, "db-host", "", "Database host")
	flag.IntVar(&cfg.dbPort, "db-port", 5432, "Database port")
	flag.StringVar(&cfg.dbName, "db-name", "", "Database name")
	flag.StringVar(&cfg.dbUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.dbPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.dbSchema, "db-schema", "public", "Database schema")
	flag.Parse()
	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP server: '%s'", cfg.HTTPPort)
	}
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s "+
		"user=%s password=%s sslmode=disable search_path=%s",
		cfg.dbHost, cfg.dbPort, cfg.dbName, cfg.dbUser, cfg.dbPassword, cfg.dbSchema)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()
	v1API := v1.NewTraineeServiceServer(db)
	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
