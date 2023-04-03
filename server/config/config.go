package config

// Config 系统配置
type Config struct {
	Service    Service    `mapstructure:"service"`
	Log        Log        `mapstructure:"log"`
	Orm        []Orm      `mapstructure:"orm"`
	Middleware Middleware `mapstructure:"middleware"`
	Redis      Redis      `mapstructure:"redis"`
	Cert       []Cert     `mapstructure:"cert"`
	Email      Email      `mapstructure:"email"`
}
