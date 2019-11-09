package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	Routes "github.com/evscott/z3-12c-api/Routes"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	conf := GetConfig(ctx, mux.NewRouter())
	_ = Routes.GetRoutes(conf.Router, conf.JenkinsClient, conf.GithubClient)

	/***** Start up *****/
	go func() {
		log.Printf("Starting up...\n")
		if err := conf.Server.ListenAndServe(); err == nil {
			log.Printf("Listening on %s...\n", conf.Server.Addr)
		} else {
			log.Fatal(err)
		}
	}()

	/***** Wait for shutdown *****/
	func(srv *http.Server) {
		interruptChan := make(chan os.Signal, 1)
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-interruptChan

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		log.Println("Shutting down...")
		os.Exit(0)
	}(conf.Server)
}
