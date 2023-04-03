package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/limeschool/easy-admin/server/core/runtime"
	"github.com/limeschool/easy-admin/server/core/trace"
	"github.com/limeschool/easy-admin/server/global"
	"go.uber.org/zap"
	"strings"
	"time"
)

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func RequestLog() gin.HandlerFunc {
	white := global.Config.Middleware.RequestLog.Whitelist

	return func(ctx *gin.Context) {
		method := strings.ToLower(ctx.Request.Method)
		path := ctx.Request.URL.String()
		if strings.Contains(path, "?") {
			path = strings.Split(path, "?")[0]
		}

		// 判断白名单
		if white[method+":"+path] {
			return
		}

		now := time.Now()
		blw := ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		req := runtime.RequestInfo(ctx)

		ctx.Next()

		end := time.Now()
		res := any(nil)
		if strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/json") {
			m := map[string]any{}
			_ = json.Unmarshal(blw.Body.Bytes(), &m)
			res = m
		} else {
			res = blw.Body.String()
		}

		trace.Logger(ctx).WithOptions(zap.WithCaller(false)).Info("request-info",
			zap.Any("path", ctx.Request.URL.Path),
			zap.Any("method", ctx.Request.Method),
			zap.Any("user_agent", ctx.Request.Header.Get("User-Agent")),
			zap.Any("timestamp", fmt.Sprintf("%vms", end.Sub(now).Milliseconds())),
			zap.Any("req", req),
			zap.Any("res", res),
			zap.Any("status", ctx.Writer.Status()),
		)
	}
}
