package configs

func (p *Pgsql) Dsn() string {
	return "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port
}

type Server struct {
	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Cors  CORS  `mapstructure:"cors" json:"cors" yaml:"cors"`
}
