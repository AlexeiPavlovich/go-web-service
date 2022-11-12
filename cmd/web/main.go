package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexei/hello-world/pkg/config"
)

const portNumber = ":8080"

func main() {
	run()
}

func run() error {
	var app config.AppConfig
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routs(&app),
	}
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	var err = srv.ListenAndServe()
	log.Fatal(err)
	return err
}
