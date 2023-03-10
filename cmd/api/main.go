package main

import (
	"context"
	"database/sql"

	"flag"
	"os"
	"time"

	"github.com/gopheramit/greenlightAPI/internal/data"
	"github.com/gopheramit/greenlightAPI/internal/jsonlog"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		rps    float64
		burst  int
		enable bool
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "api server port")
	flag.StringVar(&cfg.env, "enc", "development", "Envrionment (developmetn|staging|production")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable", "postgress")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open connections", 25, "postgres maximum open connection")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle connections", 25, "postgres maximum idle connection")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-time ", "15m", "postgres connection maximum idle time")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "rate limiter maximum request per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enable, "limiter-enable", true, "Enable rate limiter")
	flag.Parse()
	//logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}
	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)
	err = app.server()

	if err != nil {
		logger.PrintFatal(err, nil)

	}
	logger.PrintFatal(err, nil)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil

}
