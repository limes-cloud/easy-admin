package config

type Log struct {
	Level  int8
	Header string
	Field  string
	Output []string
	File   struct {
		Name      string
		MaxSize   int
		MaxBackup int
		MaxAge    int
		Compress  bool
	}
}
