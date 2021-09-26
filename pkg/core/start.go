package core

import (
	"flag"

	"log"
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

	if err = initLogger(); err != nil {
		log.Fatalf("initLogger err: %v", err)
	}

}

var (
	port    string
	runMode string
)

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.Parse()

	return nil
}
func initSetting() (e error) {
	//path, _ := os.Getwd()
	//config := path + "/configs"
	//s, e := setting.NewSetting(config, runMode, "yaml")
	//if e != nil {
	//	return
	//}
	//e = s.ReadSection("Server", &global.ServerSetting)
	//if e != nil {
	//	return
	//}
	//e = s.ReadSection("App", &global.AppSetting)
	//if e != nil {
	//	return
	//}
	//
	//if e = s.ReadSection("Database", &global.DatabaseSetting); e != nil {
	//	return
	//}
	//
	//if e = s.ReadSection("Redis", &global.RedisSetting); e != nil {
	//	return
	//}
	//
	//if e = s.ReadSection("JWT", &global.JWTSetting); e != nil {
	//	return
	//}
	//
	//if e = s.ReadSection("Log", &global.LogSetting); e != nil {
	//	return
	//}
	//
	//if e = s.ReadSection("Email", &global.EmailSetting); e != nil {
	//	return
	//}
	//
	//global.AppSetting.DefaultContextTimeout *= time.Second
	//global.JWTSetting.Expire *= time.Second
	//global.ServerSetting.ReadTimeout *= time.Second
	//global.ServerSetting.WriteTimeout *= time.Second
	//if port != "" {
	//	global.ServerSetting.HttpPort = port
	//}
	//if runMode != "" {
	//	global.ServerSetting.RunMode = runMode
	//}

	return
}
func initLogger() (e error) {

	//logSet := global.LogSetting
	//global.Logger, e = logger.NewLogger(logSet.Formatter, logSet.Level, logSet.ReportCaller, logSet.SavePath)
	//if e != nil {
	//	return
	//}

	//logger.Log.AddHook(&logger.AppHook{})

	return
}
