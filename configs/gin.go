package configs

type Gin struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Cert string `mapstructure:"cert" json:"cert" yaml:"cert"`
	Key  string `mapstructure:"key" json:"key" yaml:"key"`
}
