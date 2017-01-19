package main
import (
    "net/http"
    "log"
    "net"
)

func redirect(w http.ResponseWriter, req *http.Request) {
    host, _, _ := net.SplitHostPort(req.Host)

    target := "https://" + host + req.URL.Path
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    log.Printf("redirect to: %s", target)
    http.Redirect(w, req, target, http.StatusMovedPermanently)
}

func main() {
    // redirect every http request to https
    http.ListenAndServe(":8080", http.HandlerFunc(redirect))
}
