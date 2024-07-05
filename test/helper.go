package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	code        int         // 状态码
	param       string      // 参数
	method      string      // 请求类型
	desc        string      // 描述
	showBody    bool        // 是否展示返回
	errMsg      string      // 错误信息
	url         string      // 链接
	contentType string      //
	ext1        interface{} // 自定义1
	ext2        interface{} // 自定义2
}

func NewBufferString(body string) io.Reader {
	return bytes.NewBufferString(body)
}

func PerformRequest(mothod, url, contentType string, body string) (r *http.Request, w *httptest.ResponseRecorder) {
	return
}
