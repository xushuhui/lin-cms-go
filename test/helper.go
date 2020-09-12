package test

import (
	"encoding/json"
	"io"
	"lin-cms-go/internal/router"
	"lin-cms-go/pkg/core"
	"net/http"
	"net/http/httptest"

	"bytes"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	code        int         //状态码
	param       string      //参数
	method      string      //请求类型
	desc        string      //描述
	showBody    bool        //是否展示返回
	errMsg      string      //错误信息
	url         string      //链接
	contentType string      //
	ext1        interface{} //自定义1
	ext2        interface{} //自定义2
}

func NewBufferString(body string) io.Reader {
	return bytes.NewBufferString(body)
}

func PerformRequest(mothod, url, contentType string, body string) (c *gin.Context, r *http.Request, w *httptest.ResponseRecorder) {
	rout := router.InitRouter()

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	r = httptest.NewRequest(mothod, url, NewBufferString(body))
	c.Request = r
	c.Request.Header.Set("Content-Type", contentType)
	rout.ServeHTTP(w, r)
	return
}

func call(t *testing.T, testcase []TestCase) {
	for k, v := range testcase {
		_, _, w := PerformRequest(v.method, v.url, v.contentType, v.param)
		//assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
		fmt.Println()
		fmt.Printf("第%d个测试用例：%s", k+1, v.desc)
		if v.showBody {
			fmt.Printf("接口返回%s", w.Body.String())
			fmt.Println()
		}
		s := core.Error{}
		err := json.Unmarshal([]byte(w.Body.String()), &s)
		assert.NoError(t, err)
		assert.Equal(t, v.errMsg, s.Message, "错误信息不一致")
		assert.Equal(t, v.code, s.Code, "错误码不一致")
	}
}
