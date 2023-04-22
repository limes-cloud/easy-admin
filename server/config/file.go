package config

type File struct {
	SubDir map[string]struct {
		Accepts []string
		MaxSize int
		Rename  bool
	}
	LocalDir        string
	DriveType       string
	Endpoint        string
	SecretID        string
	SecretKey       string
	Bucket          string
	Region          string
	AccessKeyID     string
	AccessKeySecret string
	AccessKey       string
	Domain          string
	Private         bool
	Location        string
	UseSsl          bool
	Zone            string
}
