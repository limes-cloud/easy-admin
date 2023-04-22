package address

import (
	"bytes"
	"encoding/json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type address struct {
	ip string
}

type Address interface {
	GetAddress() string
}

func New(ip string) Address {
	return &address{
		ip: ip,
	}
}

func (a address) GetAddress() string {
	if a.ip == "127.0.0.1" || a.ip == "::1" {
		return "本地登陆"
	}
	if a.check() {
		// ip转地址
		if res := IPWhois(a.ip); res != "" {
			return res
		}
		return "地址查询失败"
	}
	return "非法ip地址"
}

func (a address) check() bool {
	addr := strings.Trim(a.ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func IPWhois(ip string) string {
	type response struct {
		Addr string `json:"addr"`
	}
	var resp response
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	_ = Get(url, &resp, true)
	return resp.Addr
}

func Get(url string, dst interface{}, toUtf8 bool) error {
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if toUtf8 {
		body, _ = GbkToUtf8(body)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
