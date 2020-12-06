package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/hardstylez72/bblog/ad/pkg/group"
	grouproute "github.com/hardstylez72/bblog/ad/pkg/group/route"
	"github.com/hardstylez72/bblog/ad/pkg/logger"
	"github.com/hardstylez72/bblog/ad/pkg/route"
	"github.com/hardstylez72/bblog/ad/pkg/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type Server struct {
	log    *zap.SugaredLogger
	router chi.Router
}

func main() {

	log, err := logger.New("")
	errCheck(err, "can't load config")
	defer log.Sync()

	err = NewServer(log).Run()
	errCheck(err, "can't run server")
}

func errCheck(err error, errorText string) {
	if err == nil {
		return
	}
	log.Fatal(errorText, ": ", err)
}

func NewServer(log *zap.SugaredLogger) *Server {
	return &Server{
		router: chi.NewRouter(),
		log:    log,
	}
}

func (s *Server) Run() error {

	configPath := flag.String("config", "/home/hs/go/src/github.com/hardstylez72/bblog/ad/cmd/server/config.yaml", "path to config file")
	flag.Parse()

	err := Load(*configPath)
	if err != nil {
		return err
	}
	//err = tracer.New(viper.GetString("tracer.jaeger.collectorEndpoint"), viper.GetString("tracer.jaeger.serviceName"))
	//if err != nil {
	//	return err
	//}

	httpServer := &http.Server{
		Addr:    viper.GetString("port"),
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
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	r.Use(c)

	r.Use(middleware.RequestID)
	r.Use(logger.Inject(s.log))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount(apiPathPrefix, r)

	err := Start(r)
	if err != nil {
		log.Fatal(err)
	}
	s.log.Info("app is successfully running")

	return r
}

func Start(r chi.Router) error {
	pg, err := storage.NewPGConnection(viper.GetString("databases.postgres"))
	if err != nil {
		return err
	}

	pgx, err := storage.WrapPgConnWithSqlx(pg)
	if err != nil {
		return err
	}

	var groupRepository group.Repository
	var routeRepository route.Repository
	var groupRouteRepository grouproute.Repository
	{
		groupRepository = group.NewRepository(pgx)
		routeRepository = route.NewRepository(pgx)
		groupRouteRepository = grouproute.NewRepository(pgx)
	}

	group.NewController(groupRepository).Mount(r)
	route.NewController(routeRepository).Mount(r)
	grouproute.NewController(groupRouteRepository).Mount(r)

	return nil
}
