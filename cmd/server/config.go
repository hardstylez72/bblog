package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/hardstylez72/bblog/internal/api/controller/auth"
	"github.com/hardstylez72/bblog/internal/objectstorage"
	"github.com/hardstylez72/bblog/internal/storage"
	"github.com/hardstylez72/bblog/internal/tracer"
)

const (
	cfgAllData = ""
)

type Config struct {
	Port      string
	Env       string
	Host      string
	Tracer    tracer.Config
	Oauth     auth.Oauth
	Databases Databases
	ObjectStorage
}

type ObjectStorage struct {
	Minio objectstorage.Config
}

type Databases struct {
	Postgres storage.PostgresConnect
}

func LoadFromFile(filePath string) (*Config, error) {
	config.WithOptions(func(options *config.Options) {
		options = &config.Options{
			ParseEnv:    false,
			Readonly:    true,
			EnableCache: false,
			ParseKey:    true,
			Delimiter:   0,
			DumpFormat:  "",
			ReadFormat:  "",
		}
	})
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles(filePath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = config.BindStruct(cfgAllData, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
