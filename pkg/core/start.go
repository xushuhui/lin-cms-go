package core

import (
	"flag"
	"github.com/gin-gonic/gin"
	"lin-cms-go/global"
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

	isVersion bool
)

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", gin.Mode(), "启动模式")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}
func initSetting() error {
	path, _ := os.Getwd()
	config := path + "/configs"
	s, err := setting.NewSetting(config, runMode, "yaml")
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	if err = s.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	if err = s.ReadSection("Redis", &global.RedisSetting); err != nil {
		return err
	}

	if err = s.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}

	if err = s.ReadSection("Log", &global.LogSetting); err != nil {
		return err
	}

	if err = s.ReadSection("Email", &global.EmailSetting); err != nil {
		return err
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

	return nil
}
func initLogger() (e error) {
	//global.Logger = logger.NewLogger(&lumberjack.Logger{
	//	Filename:  fileName,
	//	MaxSize:   500,
	//	MaxAge:    10,
	//	LocalTime: true,
	//}, "", log.LstdFlags).WithCaller(2)
	logSet := global.LogSetting
	global.Logger, e = logger.NewLogger(logSet.Formatter, logSet.Level, logSet.ReportCaller, logSet.SavePath)
	if e != nil {
		return
	}

	//logger.Log.AddHook(&logger.AppHook{})

	return
}

func initDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
func initTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("s", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
