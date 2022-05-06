package configs

import "github.com/go-redis/redis/v8"

func (p *Pgsql) Dsn() string {
	return "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port
}

type Server struct {
	Gin   Gin            `mapstructure:"gin" json:"gin" yaml:"gin"`
	Pgsql Pgsql          `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	JWT   JWT            `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Cors  CORS           `mapstructure:"cors" json:"cors" yaml:"cors"`
	Redis *redis.Options `mapstructure:"redis" json:"redis" yaml:"redis"`
}
