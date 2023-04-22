package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"io"
	"runtime"
	"strings"
	"unsafe"
)

// RequestInfo 获取请求参数
func RequestInfo(c *gin.Context) map[string]any {
	// 获取body
	getBody := func(c *gin.Context) any {
		// 处理form/data 上传
		if strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
			return getFormData(c)
		}

		// 读取数据
		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

		// 转换格式
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			m := map[string]any{}
			_ = json.Unmarshal(data, &m)
			return m
		}

		// 返回字符串
		return *(*string)(unsafe.Pointer(&data))
	}

	getParams := func(c *gin.Context) string {
		return c.Request.URL.Query().Encode()
	}

	return map[string]any{
		"params": getParams(c),
		"body":   getBody(c),
	}
}

// 将文件内容转换成文件名
func getFormData(ctx *gin.Context) map[string]any {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil
	}

	// 处理form-data 参数
	body := make(map[string]any)
	for key, val := range form.Value {
		if len(val) == 1 {
			body[key] = val[0]
		} else if len(val) > 1 {
			body[key] = val
		}
	}

	// 处理上传文件
	for key, files := range form.File {
		var fileNames []string
		for _, file := range files {
			fileNames = append(fileNames, file.Filename)
		}
		if len(fileNames) == 1 {
			body[key] = fileNames[0]
		} else if len(fileNames) > 1 {
			body[key] = fileNames
		}
	}
	return body
}

// PanicErr 获取panic错误堆栈
func PanicErr() []string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller
	var arr []string
	for _, pc := range pcs[:n-4] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		arr = append(arr, fmt.Sprintf("%s:%d", file, line))
	}
	return arr
}
