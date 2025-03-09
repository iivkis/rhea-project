package config

import (
	"sync"
)

type AppConfig struct {
	PgConfig PgConfig
}

var cfg *AppConfig
var cfgOnce sync.Once

func Get() *AppConfig {
	cfgOnce.Do(func() {
		cfg = &AppConfig{
			PgConfig: loadPgCfg(),
		}
	})
	return cfg
}
