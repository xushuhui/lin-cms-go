package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"lin-cms-go/global"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// AES 对称加密
func Sign() gin.HandlerFunc {

	return func(c *gin.Context) {

		sign, err := verifySign(c)

		if sign != nil {
			err = core.NewInvalidParamsError("Debug Sign")
			return
		}

		if err != nil {
			err = core.NewInvalidParamsError(err.Error())
			return
		}

		c.Next()
	}
}

// 验证签名
func verifySign(c *gin.Context) (map[string]string, error) {
	_ = c.Request.ParseForm()
	req := c.Request.Form
	debug := strings.Join(c.Request.Form["debug"], "")
	ak := strings.Join(c.Request.Form["ak"], "")
	sn := strings.Join(c.Request.Form["sn"], "")
	ts := strings.Join(c.Request.Form["ts"], "")
	var appSecret string
	// 验证来源
	value, ok := global.ApiAuthConfig[ak]
	if ok {
		appSecret = value["aes"]
	} else {
		return nil, errors.New("ak Error")
	}

	if debug == "1" {
		currentUnix := utils.GetCurrentUnix()
		req.Set("ts", strconv.FormatInt(currentUnix, 10))

		sn, err := createSign(req, appSecret)
		if err != nil {
			return nil, errors.New("sn Exception")
		}

		res := map[string]string{
			"ts": strconv.FormatInt(currentUnix, 10),
			"sn": sn,
		}
		return res, nil
	}

	// 验证过期时间
	timestamp := time.Now().Unix()
	exp, _ := strconv.ParseInt(global.AppSignExpiry, 10, 64)
	tsInt, _ := strconv.ParseInt(ts, 10, 64)
	if tsInt > timestamp || timestamp-tsInt >= exp {
		return nil, errors.New("ts Error")
	}

	// 验证签名
	if sn == "" {
		return nil, errors.New("sn Error")
	}

	decryptStr, decryptErr := utils.AesDecrypt(sn, []byte(appSecret), appSecret)
	if decryptErr != nil {
		return nil, errors.New(decryptErr.Error())
	}
	if decryptStr != createEncryptStr(req) {
		return nil, errors.New("sn Error")
	}
	return nil, nil
}

// 创建签名
func createSign(params url.Values, appSecret string) (s string, err error) {
	s, err = utils.AesEncrypt(createEncryptStr(params), []byte(appSecret), appSecret)
	if err != nil {
		return "", err
	}
	return
}

func createEncryptStr(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return str
}
