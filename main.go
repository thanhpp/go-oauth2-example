package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thanhpp/go-oauth2-example/infrastructure/apdaters/httpserver"
)

func main() {
	var (
		host = "127.0.0.1"
		port = "8000"
	)

	start, stop := httpserver.StartHTTPServer(host, port)

	log.Println("create http server OK")

	if err := start(); err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-sigs
	stop()
}
