package config

type Cert struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}
