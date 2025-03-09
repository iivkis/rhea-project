package config

import (
	"sync"

	"github.com/spf13/viper"
)

type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type AppConfig struct {
	PgConfig PgConfig
}

var cfg *AppConfig
var cfgOnce sync.Once

func Cfg() *AppConfig {
	cfgOnce.Do(func() {
		cfg = &AppConfig{
			PgConfig: loadPgCfg(),
		}
	})
	return cfg
}
func loadPgCfg() PgConfig {
	var pgCfg PgConfig

	viper.SetEnvPrefix("pg")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(pgCfg); err != nil {
		panic(err)
	}

	return pgCfg
}
