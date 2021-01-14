package main

import (
	"context"
	"fmt"
	"github/eddiefisher/gokit-account/account"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jackc/pgx/v4"
)

func main() {
	// logger settings
	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "account",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	_ = level.Info(logger).Log("msg", "service started")
	ended := func() { _ = level.Info(logger).Log("msg", "service ended") }
	defer ended()

	// database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// set channel for good exit and return errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// set repository
	repository := account.NewRepo(conn, logger)
	// initialize service
	srv := account.NewService(repository, logger)
	// set endpoints
	endpoints := account.MakeEndpoints(srv)

	// create http server
	var httpAddr = os.Getenv("HTTP_ADDR")
	var ctx = context.Background()
	go func() {
		fmt.Println("listen on port", httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(httpAddr, handler)
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
