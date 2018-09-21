package httpserver

import (
	"boiler-plate/internal/httpserver/router"
	"fmt"
	"log"
	"net/http"
)

// HTTPServer ...
type HTTPServer struct {
}

// Serve ...
func (hs *HTTPServer) Serve(port string) {
	r := router.NewV1Router()
	log.Printf("About to listen on %s. Go to http://127.0.0.1:%s", port, port)
	httpConfig := fmt.Sprintf(":%s", port)
	http.ListenAndServe(httpConfig, r)
}

// NewHTTPServer ...
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}
