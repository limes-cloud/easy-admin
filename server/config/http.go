package config

import "time"

type Http struct {
	EnableLog        bool
	RetryCount       int
	RetryWaitTime    time.Duration
	MaxRetryWaitTime time.Duration
	Timeout          time.Duration
}
