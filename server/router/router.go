package router

import (
	"net/http"

	"github.com/aixoio/bridge/server/env"
)

func StartServer(dat env.Env) {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	server := http.Server{
		Addr:    ":" + dat.Port,
		Handler: r,
	}

	server.ListenAndServe()

}
