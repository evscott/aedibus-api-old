package main

import (
	"context"
	"github.com/evscott/z3-e2c-api/shared"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	Router "github.com/evscott/z3-e2c-api/router"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	conf := shared.GetConfig(ctx, mux.NewRouter())
	Router.Init(conf.Router, conf.GithubClient, conf.DAL, conf.Logger)

	// Start up
	go func() {
		log.Printf("Starting up...\n")
		if err := conf.Server.ListenAndServe(); err == nil {
			log.Printf("Listening on %s...\n", conf.Server.Addr)
		} else {
			log.Fatal(err)
		}
	}()

	// Wait for shutdown
	func(srv *http.Server) {
		interruptChan := make(chan os.Signal, 1)
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-interruptChan

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		log.Println("Shutting down...")

		// Shutdown server
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		conf.DAL.Shutdown()

		os.Exit(0)
	}(conf.Server)
}
