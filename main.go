package main

import (
	"net/http"
	"log"
	"net"
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
	http.ListenAndServe(":8080", nil)
}
