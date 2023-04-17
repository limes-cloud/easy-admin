package config

import "time"

type Middleware struct {
	Pprof struct {
		Enable bool   `mapstructure:"enable"`
		Secret string `mapstructure:"secret"`
		Query  string `mapstructure:"query"`
	} `mapstructure:"pprof"`

	RateLimit struct {
		Enable bool `mapstructure:"enable"`
		Limit  int  `mapstructure:"limit"`
	} `mapstructure:"rateLimit"`

	IpLimit struct {
		Enable bool          `mapstructure:"enable"`
		Redis  string        `mapstructure:"redis"`
		Limit  int           `mapstructure:"limit"`
		Window time.Duration `mapstructure:"window"`
	} `mapstructure:"ipLimit"`

	CupLoadShedding struct {
		Enable    bool          `mapstructure:"enable"`
		Threshold int           `mapstructure:"threshold"`
		Bucket    int           `mapstructure:"bucket"`
		Window    time.Duration `mapstructure:"window"`
	} `mapstructure:"cupLoadShedding"`

	Cors struct {
		Enable       bool   `mapstructure:"enable"`
		AllowHeader  string `mapstructure:"allowHeader"`
		AllowMethod  string `mapstructure:"allowMethod"`
		AllowOrigin  string `mapstructure:"allowOrigin"`
		ExposeHeader string `mapstructure:"exposeHeader"`
		Credentials  bool   `mapstructure:"credentials"`
	} `mapstructure:"cors"`

	Jwt struct {
		Enable    bool            `mapstructure:"enable"`
		Redis     string          `mapstructure:"redis"`
		Header    string          `mapstructure:"header"`
		Secret    string          `mapstructure:"secret"`
		Expire    time.Duration   `mapstructure:"expire"`
		Renewal   time.Duration   `mapstructure:"renewal"`
		Unique    bool            `mapstructure:"unique"`
		Whitelist map[string]bool `mapstructure:"whitelist"`
	} `mapstructure:"jwt"`

	Casbin struct {
		Enable    bool            `mapstructure:"enable"`
		DB        string          `mapstructure:"db"`
		Whitelist map[string]bool `mapstructure:"whitelist"`
	} `mapstructure:"casbin"`

	RequestLog struct {
		Enable    bool            `mapstructure:"enable"`
		Whitelist map[string]bool `mapstructure:"whitelist"`
	} `mapstructure:"requestLog"`
}
