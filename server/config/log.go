package config

type Log struct {
	Level  int8     `mapstructure:"level"` // 日志等级
	Header string   `json:"header"`
	Output []string `mapstructure:"output"` // 输出方式 ["stdout","file"]
	File   struct {
		Name      string `mapstructure:"name"`      // 输出文件名
		MaxSize   int    `mapstructure:"maxSize"`   // 文件最大容量 单位M
		MaxBackup int    `mapstructure:"maxBackup"` // 最大备份
		MaxAge    int    `mapstructure:"maxAge"`    // 最长存储时间
		Compress  bool   `mapstructure:"compres"`   // 是否压缩归档
	} `mapstructure:"file"`
}
