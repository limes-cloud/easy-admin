package config

import "time"

type Orm struct {
	Enable        bool
	Drive         string
	Name          string
	Dsn           string
	MaxLifetime   time.Duration
	MaxOpenConn   int
	MaxIdleConn   int
	Level         int
	SlowThreshold time.Duration
	Replicas      []string
}
