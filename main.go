package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"urlShorter/config"
)

func main() {
	var conf Config
	config.ReadConfig(&conf, "params/config.yaml")
	srv := NewServer(conf)
	// Start the server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
		}
	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
