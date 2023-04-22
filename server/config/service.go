package config

import "time"

type Service struct {
	Name        string
	Timeout     time.Duration
	Addr        string
	Debug       bool
	ErrorCode   int
	SuccessCode int
}
