package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/hardstylez72/bbckend/internal/controller"
	"github.com/hardstylez72/bbckend/internal/store"
	"github.com/hardstylez72/bbckend/internal/tracer"
)

const (
	cfgAllData = ""
)

type Config struct {
	Port   string
	Env    string
	Host   string
	Tracer tracer.Config
	Oauth  controller.Oauth
	Databases Databases
}

type Databases struct {
	Postgres store.Postgres
}

func Load(filePath string) (*Config, error) {
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
