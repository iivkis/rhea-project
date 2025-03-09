package config

import "github.com/spf13/viper"

type PgConfigDriver string

type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c *PgConfig) GetDSN() string {
	return "postgres://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=disable"
}

func loadPgCfg() PgConfig {
	var pgCfg PgConfig

	viper.AutomaticEnv()

	viper.BindEnv("host", "pg_host")
	viper.BindEnv("port", "pg_port")
	viper.BindEnv("user", "pg_user")
	viper.BindEnv("password", "pg_pass")
	viper.BindEnv("database", "pg_name")

	if err := viper.Unmarshal(&pgCfg); err != nil {
		panic(err)
	}

	return pgCfg
}
