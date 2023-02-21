package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"memoriesbox/pkg/db"
	"memoriesbox/pkg/logger"
	"memoriesbox/pkg/repositories"
	"memoriesbox/pkg/web"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/peterbourgon/ff/v3"
)

//go:embed templates static
var content embed.FS

func main() {
	fs := flag.NewFlagSet("memoriesbox", flag.ContinueOnError)

	dbConfig := db.GetConfig(fs)
	webConfig := web.GetConfig(fs)
	repoConfig := repositories.GetConfig(fs)

	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("MEMORIESBOX"),
	)
	if err != nil {
		log.Fatal(err)
	}

	loggerApp := logger.New(fs)
	dbApp := db.New(dbConfig, loggerApp)
	repoApp := repositories.New(repoConfig, loggerApp, dbApp)
	webApp := web.New(webConfig, content, loggerApp, repoApp)

	go webApp.Start()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := webApp.Shutdown(ctx); err != nil {
		loggerApp.Sugar.Error(err)
	}
}
