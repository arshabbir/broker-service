package api

import (
	"net/http"

	"github.com/arshabbir/brokermod/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type server struct {
	cfg *config.Config
}

type Server interface {
	Start() error
}

func NewServer(cfg *config.Config) Server {
	return &server{cfg: cfg}
}
func (s *server) Start() error {
	//  Add routes

	mux := routes(s)
	srv := http.Server{
		Addr:    s.cfg.AppPort,
		Handler: mux,
	}

	srv.ListenAndServe()

	return nil

}
func routes(s *server) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	r.Get("/ping", s.HandlePing)
	r.Post("/", s.Broker)

	r.Post("/auth", s.HandleAuth)

	return r
}
