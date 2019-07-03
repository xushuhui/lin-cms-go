/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package tests

import (
	"fmt"
	"github.com/astaxie/beego"
	"lin-cms-beego/core"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	fmt.Println(apppath + "/conf/app.conf")
	beego.TestBeegoInit(apppath)
	beego.LoadAppConfig("ini", apppath+"/conf/app.conf")

}
func TestGetBooks(t *testing.T) {
	core.InitEnv()
	r, _ := http.NewRequest("GET", "/v1/books", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace(w.Code, w.Body.String())

	//Convey("Subject: Test Station Endpoint\n", t, func() {
	//        Convey("Status Code Should Be 200", func() {
	//                So(w.Code, ShouldEqual, 200)
	//        })
	//        Convey("The Result Should Not Be Empty", func() {
	//                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	//        })
	//})
}
