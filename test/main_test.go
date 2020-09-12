package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func setup() {

	gin.SetMode(gin.TestMode)
	//fmt.Println(global.JWTSetting.Secret)
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}
func TestMain(m *testing.M) {
	setup()
	fmt.Println("Test begins....")
	code := m.Run() // 如果不加这句，只会执行Main
	teardown()
	os.Exit(code)
}
