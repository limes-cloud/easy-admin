package config

type Enforcer struct {
	Enable    bool
	DB        string
	Whitelist map[string]bool
}
