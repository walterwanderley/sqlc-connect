// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect).

package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// database driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:generate sqlc-connect -m booktest -append

const serviceName = "booktest"

var (
	dbURL string
	port  int
)

func main() {
	var dev bool
	flag.StringVar(&dbURL, "db", "", "The Database connection URL")
	flag.IntVar(&port, "port", 5000, "The server port")
	flag.BoolVar(&dev, "dev", false, "Set logger to development mode")
	flag.Parse()

	log := logger(dev)
	defer log.Sync()

	if err := run(log); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("server error", zap.Error(err))
		os.Exit(1)
	}
}

func run(log *zap.Logger) error {
	if _, err := maxprocs.Set(); err != nil {
		log.Warn("startup", zap.Error(err))
	}
	log.Info("startup", zap.Int("GOMAXPROCS", runtime.GOMAXPROCS(0)))

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	registerHandlers(mux, log, db)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: h2c.NewHandler(mux, &http2.Server{}),
		// Please, configure timeouts!
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-done
		log.Warn("signal detected...", zap.Any("signal", sig))
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	log.Info("Listening...", zap.Int("port", port))
	return server.ListenAndServe()
}

func logger(dev bool) *zap.Logger {
	var config zap.Config
	if dev {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": serviceName,
	}

	log, err := config.Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return log
}
