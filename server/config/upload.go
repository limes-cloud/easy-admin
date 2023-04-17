package config

type Upload struct {
	AcceptTypes     []string `json:"accept_types" mapstructure:"accept_types"`
	MaxSize         int      `json:"max_size" mapstructure:"max_size"`
	Rename          bool     `json:"rename" mapstructure:"rename"`
	LocalDir        string   `json:"local_dir" mapstructure:"local_dir"`
	DriveType       string   `json:"drive_type" mapstructure:"drive_type"`
	Endpoint        string   `json:"endpoint" mapstructure:"endpoint"`
	SecretID        string   `json:"secret_id" mapstructure:"secret_id"`
	SecretKey       string   `json:"secret_key" mapstructure:"secret_key"`
	Bucket          string   `json:"bucket" mapstructure:"bucket"`
	Region          string   `json:"region" mapstructure:"region"`
	AccessKeyID     string   `json:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string   `json:"access_key_secret" mapstructure:"access_key_secret"`
	AccessKey       string   `json:"access_key" mapstructure:"access_key"`
	Domain          string   `json:"domain" mapstructure:"domain"`
	Private         bool     `json:"private" mapstructure:"private"`
	Location        string   `json:"location" mapstructure:"location"`
	UseSsl          bool     `json:"use_ssl" mapstructure:"use_ssl"`
	Zone            string   `json:"zone" mapstructure:"zone"`
}
