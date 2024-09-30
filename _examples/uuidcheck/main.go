// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect).

package main

import (
	"context"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"github.com/flowchartsman/swaggerui"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// database driver
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate /var/folders/ks/_614mn9d3c53zg_y00d9md840000gn/T/go-build1798372745/b001/exe/sqlc-connect -append

const serviceName = "uuidcheck"

var (
	dbURL string
	port  int

	//go:embed api/apidocs.swagger.json
	openAPISpec []byte
)

func main() {
	var dev bool
	flag.StringVar(&dbURL, "db", "", "The Database connection URL")
	flag.IntVar(&port, "port", 5000, "The server port")

	flag.BoolVar(&dev, "dev", false, "Set logger to development mode")

	flag.Parse()

	initLogger(dev)

	if err := run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	_, err := maxprocs.Set()
	if err != nil {
		slog.Warn("startup", "error", err)
	}
	slog.Info("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	mux := http.NewServeMux()
	var interceptors []connect.Interceptor

	registerHandlers(mux, db, interceptors)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(openAPISpec)))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: h2c.NewHandler(mux, &http2.Server{}),
		// Please, configure timeouts!
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-done
		slog.Warn("signal detected...", "signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	slog.Info("Listening...", "port", port)
	return server.ListenAndServe()
}

func initLogger(dev bool) {
	var handler slog.Handler
	opts := slog.HandlerOptions{
		AddSource: true,
	}
	switch {
	case dev:
		handler = slog.NewTextHandler(os.Stderr, &opts)
	default:
		handler = slog.NewJSONHandler(os.Stderr, &opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
