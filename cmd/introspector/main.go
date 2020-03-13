package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-http-utils/logger"
	"github.com/namsral/flag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "Set the port to listen on")
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("OK"))
	})

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	handler := logger.Handler(mux, os.Stdout, logger.CombineLoggerType)

	// Serve h2c (HTTP/2.0 ClearText, instead of TLS HTTP2) for Envoy
	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}
