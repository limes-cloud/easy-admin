/**
 * @Author: 1280291001@qq.com
 * @Description:
 * @File: request
 * @Version: 1.0.0
 * @Date: 2023/4/21 23:38
 */

package http

import (
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
	"github.com/limeschool/easy-admin/server/config"
	"go.uber.org/zap"
	"time"
	"unsafe"
)

type request struct {
	c        *config.Http
	request  *resty.Request
	logger   *zap.Logger
	inputLog bool
}

type Request interface {
	DisableLog() Request
	Option(fn RequestFunc) Request
	Get(url string) (*response, error)
	Post(url string, data interface{}) (*response, error)
	PostJson(url string, data interface{}) (*response, error)
	Put(url string, data interface{}) (*response, error)
	PutJson(url string, data interface{}) (*response, error)
	Delete(url string) (*response, error)
}

func New(conf *config.Http, logger *zap.Logger) Request {
	client := resty.New()
	if conf.MaxRetryWaitTime == 0 {
		conf.RetryWaitTime = 5 * time.Second
	}
	if conf.Timeout == 0 {
		conf.Timeout = 60 * time.Second
	}
	client.SetRetryWaitTime(conf.RetryWaitTime)
	client.SetRetryMaxWaitTime(conf.MaxRetryWaitTime)
	client.SetRetryCount(conf.RetryCount)
	client.SetTimeout(conf.Timeout)
	return &request{
		request:  client.R(),
		logger:   logger,
		inputLog: true,
	}
}

type RequestFunc func(*resty.Request) *resty.Request

func (h *request) DisableLog() Request {
	h.inputLog = false
	return h
}

func (h *request) Option(fn RequestFunc) Request {
	h.request = fn(h.request)
	return h
}

func (h *request) log(t int64, res *response) {
	if !(h.c.EnableLog && h.inputLog) {
		return
	}

	resData := res.Body()
	logs := []zap.Field{
		zap.Any("method", h.request.Method),
		zap.Any("url", h.request.URL),
		zap.Any("header", h.request.Header),
		zap.Any("body", h.request.Body),
		zap.Any("cost", time.Now().UnixMilli()-t),
		zap.Any("res", *(*string)(unsafe.Pointer(&resData))),
	}
	if len(h.request.FormData) != 0 {
		logs = append(logs, zap.Any("form-data", h.request.FormData))
	}
	if len(h.request.QueryParam) != 0 {
		logs = append(logs, zap.Any("query", h.request.QueryParam))
	}
	h.logger.Info("request", logs...)
}

func (h *request) Get(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Get(url)
	return res, res.err
}

func (h *request) Post(url string, data interface{}) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Post(url)
	return res, res.err
}

func (h *request) PostJson(url string, data interface{}) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Post(url)
	return res, res.err
}

func (h *request) Put(url string, data interface{}) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Put(url)
	return res, res.err
}

func (h *request) PutJson(url string, data interface{}) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Put(url)
	return res, res.err
}

func (h *request) Delete(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Delete(url)
	return res, res.err
}

type response struct {
	err      error
	response *resty.Response
}

func (r *response) Body() []byte {
	return r.response.Body()
}

func (r *response) Result(val interface{}) error {
	return json.Unmarshal(r.response.Body(), val)
}
