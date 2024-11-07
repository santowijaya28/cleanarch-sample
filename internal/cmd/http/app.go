package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santowijaya28/cleanarch-sample/internal/src/config"
	"github.com/santowijaya28/cleanarch-sample/internal/src/core/postgres"
	dHttp "github.com/santowijaya28/cleanarch-sample/internal/src/delivery/http"
	rUser "github.com/santowijaya28/cleanarch-sample/internal/src/repository/user"
	"github.com/santowijaya28/cleanarch-sample/pkg/log"
)

func main() {
	// init logger
	log.Init()
	defer log.Sync()

	log.Info("Success load log...")

	// init config
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed init config %v", err)
	}

	// init db
	pqConn, err := postgres.InitConn(cfg)
	if err != nil {
		log.Fatal("failed when initialize conn", err)
	}

	// initialize repository
	// TODO: integrate with usecase
	_ = rUser.Init(pqConn.GetUserDB())

	// initialize rest delivery
	httpDelivery := dHttp.Init()

	// initialize router
	r := mux.NewRouter()
	r.HandleFunc("/ping", httpDelivery.Ping).Methods(http.MethodGet)
	r.HandleFunc("/user/register", httpDelivery.RegisterUser).Methods(http.MethodPost)

	log.Info("Starting the server...")
	err = http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
