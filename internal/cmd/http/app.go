package main

import (
	"net/http"

	"github.com/gorilla/mux"
	dHttp "github.com/santowijaya28/cleanarch-sample/internal/src/delivery/http"
	"github.com/santowijaya28/cleanarch-sample/pkg/log"
)

func main() {
	// init logger
	log.Init()
	defer log.Sync()

	log.Info("Success load log...")

	// initialize rest delivery
	httpDelivery := dHttp.Init()

	// initialize router
	r := mux.NewRouter()
	r.HandleFunc("/ping", httpDelivery.Ping).Methods(http.MethodGet)

	log.Info("Starting the server...")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
