package debug

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer(addr string) *http.Server {
	r := chi.NewRouter()

	r.Handle("/v1/internal/buildinfo", BuildInformation())
	r.Mount("/v1/internal/debug", middleware.Profiler())

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func BuildInformation() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		info, ok := debug.ReadBuildInfo()
		w.Header().Set("Content-Type", "application/json")
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"error": "unable to read build information",
			})
			return
		}
		if err := json.NewEncoder(w).Encode(info); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"error": err,
			})
		}
	})
}
