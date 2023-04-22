package config

type Email struct {
	Template []struct {
		Name    string
		Subject string
		Src     string
	}
	Company  string
	User     string
	Host     string
	Password string
}
