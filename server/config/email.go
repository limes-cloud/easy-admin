package config

type Email struct {
	Template map[string]string `mapstructure:"template"`
	Company  string            `mapstructure:"company"`
	User     string            `mapstructure:"user"`
	Host     string            `mapstructure:"host"`
	Password string            `mapstructure:"password"`
}
