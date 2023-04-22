package config

import "time"

type Captcha struct {
	Name     string
	Type     string
	Length   int
	Expire   time.Duration
	Cache    string
	Height   int
	Width    int
	Skew     float64
	DotCount int
	Template string
}
