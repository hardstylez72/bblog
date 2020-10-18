package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	article2 "github.com/hardstylez72/bblog/internal/api/controller/article"
	"github.com/hardstylez72/bblog/internal/api/controller/auth"
	objectstorage2 "github.com/hardstylez72/bblog/internal/api/controller/objectstorage"
	user2 "github.com/hardstylez72/bblog/internal/api/controller/user"
	usermw "github.com/hardstylez72/bblog/internal/auth"
	"github.com/hardstylez72/bblog/internal/logger"
	"github.com/hardstylez72/bblog/internal/objectstorage"
	"github.com/hardstylez72/bblog/internal/storage"
	"github.com/hardstylez72/bblog/internal/storage/article"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"github.com/hardstylez72/bblog/internal/tracer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type Server struct {
	log      *logrus.Logger
	router   chi.Router
	services *Services
}

type Services struct {
	userStorage    user.Storage
	articleStorage article.Storage
	objectStorage  objectstorage.Storage
}

func main() {
	configPath := flag.String("config", "cmd/server/config.yaml", "path to config file")
	flag.Parse()

	err := Load(*configPath)
	errCheck(err, "can't load config")
	err = tracer.New(viper.GetString("tracer.jaeger.collectorEndpoint"), viper.GetString("tracer.jaeger.serviceName"))
	errCheck(err, "can't load config")

	services, err := initServices()
	errCheck(err, "can't init internal services")

	err = NewServer(services).Run()
	errCheck(err, "can't run server")
}

func errCheck(err error, errorText string) {
	if err == nil {
		return
	}
	log.Fatal(errorText, ": ", err)
}

func initServices() (*Services, error) {

	pgConn, err := storage.NewPGConnection(viper.GetString("databases.postgres"))
	if err != nil {
		return nil, err
	}
	userStorage := user.NewPGStorage(pgConn)

	//resolveDefaultAdminUser()

	articleStorage := article.NewPgStorage(pgConn)

	minioClient, err := objectstorage.NewMinioClient(objectstorage.Config{
		Host:            viper.GetString("objectStorage.minio.host"),
		AccessKeyID:     viper.GetString("objectStorage.minio.accessKeyID"),
		SecretAccessKey: viper.GetString("objectStorage.minio.secretAccessKey"),
		UseSSL:          viper.GetBool("objectStorage.minio.useSSL"),
	})
	if err != nil {
		return nil, err
	}
	objectStorage := objectstorage.NewMinioStorage(minioClient)

	return &Services{
		userStorage:    userStorage,
		articleStorage: articleStorage,
		objectStorage:  objectStorage,
	}, nil
}

func resolveDefaultAdminUser(userStorage user.Storage) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	userAmount, err := userStorage.GetUsersAmount(ctx)
	if err != nil {
		return err
	}

	if userAmount == 0 {
		userStorage.SaveUser(ctx, &user.User{
			Id:    1,
			Email: sql.NullString{},
			Login: sql.NullString{
				String: "",
				Valid:  false,
			},
		})
	}

	return nil
}

func NewServer(services *Services) *Server {
	return &Server{
		router:   chi.NewRouter(),
		log:      logger.New(),
		services: services,
	}
}

func (s *Server) Run() error {
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
		AllowedOrigins:   []string{"http://localhost:*"},
		AllowCredentials: true,
		Debug:            true,
	})
	r.Use(c)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(usermw.InjectUserIdFromCookies)
	r.Mount(apiPathPrefix, r)

	s.log.Info("app is successfully running")

	oauthCfg := auth.Oauth{
		Google: auth.Config{
			RedirectURL:  viper.GetString("oauth.google.redirectURL"),
			ClientID:     viper.GetString("oauth.google.clientID"),
			ClientSecret: viper.GetString("oauth.google.clientSecret"),
			Scopes:       viper.GetStringSlice("oauth.google.scopes"),
			UserInfoURL:  viper.GetString("oauth.google.userInfoURL"),
			UserRedirects: auth.UserRedirects{
				OnSuccess: viper.GetString("oauth.google.userRedirects.onSuccess"),
				OnFailure: viper.GetString("oauth.google.userRedirects.onFailure"),
			},
		},
		Github: auth.Config{
			RedirectURL:  viper.GetString("oauth.github.redirectURL"),
			ClientID:     viper.GetString("oauth.github.clientID"),
			ClientSecret: viper.GetString("oauth.github.clientSecret"),
			Scopes:       viper.GetStringSlice("oauth.github.scopes"),
			UserInfoURL:  viper.GetString("oauth.github.userInfoURL"),
			UserRedirects: auth.UserRedirects{
				OnSuccess: viper.GetString("oauth.github.userRedirects.onSuccess"),
				OnFailure: viper.GetString("oauth.github.userRedirects.onFailure"),
			},
		},
		SessionCookieConfig: auth.SessionCookieConfig{
			Name:   viper.GetString("oauth.sessionCookie.name"),
			Domain: viper.GetString("oauth.sessionCookie.domain"),
			Path:   viper.GetString("oauth.sessionCookie.path"),
			MaxAge: viper.GetInt("oauth.sessionCookie.maxAge"),
			Secure: viper.GetBool("oauth.sessionCookie.secure"),
		},
	}

	auth.NewAuthController(oauthCfg).Mount(s.router)
	article2.NewArticleController(s.services.articleStorage).Mount(s.router)
	objectstorage2.NewObjectStorageController(s.services.objectStorage).Mount(s.router)
	user2.NewUserController(s.services.userStorage).Mount(s.router)

	return r
}
