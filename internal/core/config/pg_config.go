package config

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
