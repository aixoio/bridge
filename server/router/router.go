package router

import (
	"net/http"

	"github.com/aixoio/bridge/server/env"
	"github.com/aixoio/bridge/server/router/handlers"
)

func StartServer(dat env.Env) {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/user/signup", handlers.Signup)

	server := http.Server{
		Addr:    ":" + dat.Port,
		Handler: r,
	}

	server.ListenAndServe()

}
