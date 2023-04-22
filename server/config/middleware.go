package config

import "time"

type Middleware struct {
	Pprof struct {
		Enable bool
		Secret string
		Query  string
	}

	RateLimit struct {
		Enable bool
		Limit  int
	}

	IpLimit struct {
		Enable bool
		Cache  string
		Limit  int
		Window time.Duration
	}

	CupLoadShedding struct {
		Enable    bool
		Threshold int
		Bucket    int
		Window    time.Duration
	}

	Cors struct {
		Enable       bool
		AllowHeader  string
		AllowMethod  string
		AllowOrigin  string
		ExposeHeader string
		Credentials  bool
	}

	RequestLog struct {
		Enable    bool
		Whitelist map[string]bool
	}
}
