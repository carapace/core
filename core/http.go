package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Start the HTTP server, serving health and readiness probes.
//
// serveHTTP should be one of the last functions called.
func (a *App) serveHTTP() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", a.HealthManager.ServeHTTP)
	r.HandleFunc("/ready", a.ReadyHTTP)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", a.Health.Host, a.Health.Port), r)
	panic(err)
}

func (a *App) ReadyHTTP(w http.ResponseWriter, r *http.Request) {
	if !a.ready {
		http.Error(w, "", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
