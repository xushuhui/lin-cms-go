package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lin-cms-go/global"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		rawData, _ := c.GetRawData()

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))
		//开始时间
		startTime := utils.GetCurrentMilliUnix()
		//处理请求
		c.Next()

		//结束时间
		endTime := utils.GetCurrentMilliUnix()

		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			Error := core.Error{}
			err := json.Unmarshal([]byte(responseBody), &Error)
			if err == nil {
				responseCode = Error.Code
				responseMsg = Error.Message
				responseData = Error.Data
			}
		}

		//日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_referer"] = c.Request.Referer()

		accessLogMap["request_post_data"] = string(rawData)
		accessLogMap["request_client_ip"] = c.ClientIP()

		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"] = responseMsg
		accessLogMap["response_data"] = responseData

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)

		accessLogJson, _ := utils.JSONEncode(accessLogMap)
		AppAccessLogName := utils.LogDir(global.LogSetting.SavePath) + "access.log"
		if f, err := os.OpenFile(AppAccessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777); err != nil {
			log.Println(err)
		} else {
			f.WriteString(string(accessLogJson) + "\n")
		}
	}
}
