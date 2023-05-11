package config

type Service struct {
	Name        string `json:"name"`
	Addr        string `json:"addr"`
	Debug       bool   `json:"debug"`
	ErrorCode   int    `json:"errorCode"`
	SuccessCode int    `json:"successCode"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Copyright   string `json:"copyright"`
	StaticUrl   string `json:"staticUrl"`
	Logo        string `json:"logo"`
	TopMenu     bool   `json:"topMenu"`
}
