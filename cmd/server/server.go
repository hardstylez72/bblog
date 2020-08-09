package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	article2 "github.com/hardstylez72/bblog/internal/api/controller/article"
	"github.com/hardstylez72/bblog/internal/api/controller/auth"
	"github.com/hardstylez72/bblog/internal/logger"
	"github.com/hardstylez72/bblog/internal/storage"
	"github.com/hardstylez72/bblog/internal/storage/article"
	"github.com/hardstylez72/bblog/internal/storage/middleware2"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"github.com/hardstylez72/bblog/internal/tracer"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	log      *logrus.Logger
	config   *Config
	router   chi.Router
	services *Services
}

type Services struct {
	userStorage    user.Storage
	articleStorage article.Storage
}

func main() {
	configPath := flag.String("c", ".", "path to config file")
	flag.Parse()

	cfg, err := LoadFromFile(*configPath)
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

func initServices(cfg *Config) (*Services, error) {

	pgConn, err := storage.NewPGConnection(cfg.Databases.Postgres)
	if err != nil {
		return nil, err
	}
	userStorage := user.NewPGStorage(pgConn)
	articleStorage := article.NewPgStorage(pgConn)

	return &Services{
		userStorage:    userStorage,
		articleStorage: articleStorage,
	}, nil
}

func NewServer(config *Config, services *Services) *Server {
	return &Server{
		config:   config,
		router:   chi.NewRouter(),
		log:      logger.New(config.Env),
		services: services,
	}
}

func (s *Server) Run() error {
	httpServer := &http.Server{
		Addr:    s.config.Port,
		Handler: s.Handler(),
	}

	return httpServer.ListenAndServe()
}

func (s *Server) Handler() chi.Router {
	const (
		apiPathPrefix = "/api"
	)

	r := s.router
	c := cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	})
	r.Use(c)

	r.Use(middleware.RequestID)

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware2.LoginCheck)
	r.Mount(apiPathPrefix, r)

	s.log.Info("app is successfully running")

	auth.NewAuthController(s.config.Oauth, s.services.userStorage).Mount(s.router)
	article2.NewArticleController(s.services.articleStorage).Mount(s.router)

	return r
}
