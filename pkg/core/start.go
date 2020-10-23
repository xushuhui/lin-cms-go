package core

import (
	"flag"
	"github.com/gin-gonic/gin"
	"lin-cms-go/global"
	"lin-cms-go/internal/cache"
	"lin-cms-go/internal/model"
	"lin-cms-go/pkg/logger"
	"lin-cms-go/pkg/setting"
	"lin-cms-go/pkg/tracer"
	"log"
	"os"
	"time"
)

func StartModule() {
	//initRedis()
	var err error
	if err = initFlag(); err != nil {
		log.Fatalf("initFlag err: %v", err)
	}
	if err = initSetting(); err != nil {
		log.Fatalf("initSetting err: %v", err)
	}
	if err = initValidate(); err != nil {
		log.Fatalf("initValidate err: %v", err)
	}

	if err = initLogger(); err != nil {
		log.Fatalf("initLogger err: %v", err)
	}
	err = initDBEngine()
	if err != nil {
		log.Fatalf("initDBEngine err: %v", err)
	}
	//err = initRedis()
	//if err != nil {
	//log.Fatalf("initRedis err: %v", err)
	//}

	err = initTracer()
	if err != nil {
		log.Fatalf("initTracer err: %v", err)
	}
}

var (
	port    string
	runMode string
)

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", gin.Mode(), "启动模式")
	flag.Parse()

	return nil
}
func initSetting() (e error) {
	path, _ := os.Getwd()
	config := path + "/configs"
	s, e := setting.NewSetting(config, runMode, "yaml")
	if e != nil {
		return
	}
	e = s.ReadSection("Server", &global.ServerSetting)
	if e != nil {
		return
	}
	e = s.ReadSection("App", &global.AppSetting)
	if e != nil {
		return
	}

	if e = s.ReadSection("Database", &global.DatabaseSetting); e != nil {
		return
	}

	if e = s.ReadSection("Redis", &global.RedisSetting); e != nil {
		return
	}

	if e = s.ReadSection("JWT", &global.JWTSetting); e != nil {
		return
	}

	if e = s.ReadSection("Log", &global.LogSetting); e != nil {
		return
	}

	if e = s.ReadSection("Email", &global.EmailSetting); e != nil {
		return
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return
}
func initLogger() (e error) {

	logSet := global.LogSetting
	global.Logger, e = logger.NewLogger(logSet.Formatter, logSet.Level, logSet.ReportCaller, logSet.SavePath)
	if e != nil {
		return
	}

	//logger.Log.AddHook(&logger.AppHook{})

	return
}

func initDBEngine() (e error) {

	global.DBEngine, e = model.NewDBEngine(global.DatabaseSetting)
	if e != nil {
		return
	}

	return
}
func initTracer() (e error) {
	jaegerTracer, _, e := tracer.NewJaegerTracer("s", "127.0.0.1:6831")
	if e != nil {
		return e
	}
	global.Tracer = jaegerTracer
	return nil
}
func initRedis() (e error) {
	global.RDB, e = cache.NewRedisClient(global.RedisSetting)
	if e != nil {
		return
	}
	return
}
