package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jawher/mow.cli"
)

const (
	NAME = "uw-test"
	DESC = ""
)

func main() {
	app := cli.App(NAME, DESC)

	httpPort := app.Int(cli.IntOpt{
		Name:   "http-port",
		Desc:   "Port to listen on for HTTP connections",
		EnvVar: "HTTP_PORT",
		Value:  8080,
	})

	app.Action = func() {
		initialiseHttpServer(*httpPort)
	}
	app.Run(os.Args)
}

func initialiseHttpServer(port int) {
	router := mux.NewRouter()

	loggingHandler := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), loggingHandler))
}
