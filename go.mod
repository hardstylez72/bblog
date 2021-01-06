module github.com/hardstylez72/bblog

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.1.0
	contrib.go.opencensus.io/integrations/ocsql v0.1.4
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-chi/render v1.0.1
	github.com/go-playground/validator/v10 v10.3.0
	github.com/google/uuid v1.1.1
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.0+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/minio/minio-go/v7 v7.0.4
	github.com/pkg/errors v0.8.1
	github.com/pressly/goose v2.6.0+incompatible
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.5.1
	go.opencensus.io v0.22.1
	go.uber.org/zap v1.10.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)
