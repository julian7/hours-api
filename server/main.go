package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/app/health"

	"github.com/julian7/hours-api/handlers"
)

const version = "1.0.0"

func main() {
	var (
		httpAddr = flag.String("http", "0.0.0.0:4000", "HTTP service address.")
	)
	flag.Parse()

	log.Println("Starting server...")
	log.Printf("HTTP service listening on %s", *httpAddr)

	errChan := make(chan error, 10)

	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()

	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	httpServer.Handler = ghandlers.CORS()(handlers.LoggingHandler(router))

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Printf("Captured %v. Exiting...", s)
			health.SetReadinessStatus(http.StatusServiceUnavailable)
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
