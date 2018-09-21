package main

import (
	"boiler-plate/config"
	"boiler-plate/internal/httpserver"
)

func main() {
	server := httpserver.NewHTTPServer()
	server.Serve(config.GetConf("HTTP_PORT"))
}

