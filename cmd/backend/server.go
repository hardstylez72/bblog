package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hardstylez72/bbckend/internal/controller"
	"github.com/hardstylez72/bbckend/internal/logger"
	"github.com/hardstylez72/bbckend/internal/store"
	"github.com/hardstylez72/bbckend/internal/tracer"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	log    *logrus.Logger
	config *Config
	router chi.Router
	services *Services
}

type Services struct {
	pg store.Store
}

func (s *Server) Run() error {
	httpServer := &http.Server{
		Addr:    s.config.Port,
		Handler: s.Handler(),
	}

	return httpServer.ListenAndServe()
}

func (s *Server) Handler() chi.Router {
	r := s.router
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount("/api", r)
	s.log.Info("app is successfully running")

	r.HandleFunc("/", controller.HandleMain)

	authController, err := controller.NewAuthController(s.config.Oauth, s.services.pg)
	if err != nil {
		panic(err)
	}
	authController.Mount(s.router)

	return r
}

func NewServer(config *Config, services *Services) *Server {
	return &Server{
		config: config,
		router: chi.NewRouter(),
		log:    logger.New(config.Env),
		services: services,
	}
}

func initServices(cfg *Config) (*Services, error) {

	pg, err := store.NewPGStore(cfg.Databases.Postgres)
	if err != nil {
		return nil, err
	}

	return &Services{
		pg: pg,
	}, nil
}

func main() {
	cfg, err := Load("cmd/backend/config.example.yaml")
	errCheck(err, "can't load config")
	err = tracer.New(cfg.Tracer)
	errCheck(err, "can't load config")

	services, err := initServices(cfg)
	errCheck(err, "can't init internal services")

	err = NewServer(cfg, services).Run()
	errCheck(err, "can't run server")
}


func errCheck(err error, errorText string) {
	if err == nil {
		return
	}
	panic(err)
}


