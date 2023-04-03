package config

import "time"

type Service struct {
	Timeout time.Duration `mapstructure:"timeout"`
	Name    string        `mapstructure:"name"`
	Addr    string        `mapstructure:"addr"`
	Debug   bool          `mapstructure:"debug"`
}
