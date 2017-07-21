package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func healthcheck_handler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func redirect_handler(w http.ResponseWriter, req *http.Request) {
	host, _, err := net.SplitHostPort(req.Host)
	if err != nil {
		host = req.Host
	}

	target := "https://" + host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusMovedPermanently)
}

func main() {
	// redirect every http request to https
	http.HandleFunc("/status", healthcheck_handler)
	http.HandleFunc("/", redirect_handler)

	srv := &http.Server{
		Handler: nil,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println()
		log.Printf("Received signal: %v", sig)
		log.Println("Shutting down...")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		srv.Shutdown(ctx)
	}()

	log.Println("Listening on addr: :8080")
	log.Fatal(srv.ListenAndServe())
}
