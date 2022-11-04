package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/defer-panic/url-shortener-api/internal/config"
	"github.com/defer-panic/url-shortener-api/internal/db"
	"github.com/defer-panic/url-shortener-api/internal/server"
	"github.com/defer-panic/url-shortener-api/internal/shorten"
	"github.com/defer-panic/url-shortener-api/internal/storage"
)

func main() {
	dbCtx, dbCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer dbCancel()

	edgeClient, err := db.Connect(dbCtx, config.Get().DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	var (
		edgeStorage = storage.NewEdgeDB(edgeClient.Client())
		service     = shorten.NewService(edgeStorage)
		srv         = server.New(service)
	)

	srv.AddCloser(edgeClient.Close)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(config.Get().ListenAddr(), srv); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error running server: %v", err)
		}
	}()

	log.Println("server started")
	<-quit

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("error closing server: %v", err)
	}

	log.Println("server stopped")
}
