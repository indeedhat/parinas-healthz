package main

import (
	"net/http"

	"github.com/indeedhat/parity-nas/pkg/logging"
	"github.com/indeedhat/parity-nas/pkg/server_mux"
)

func Init(router servermux.Router, logger *logging.Logger) error {
	logger.Info("initializing healthz plugin")

	router.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		servermux.NoContent(w)
	})

	return nil
}
