package main

import (
	"flag"
	"io"
	"os"
	"time"

	"lin-cms-go/internal/conf"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	id, _    = os.Hostname()
	Name     = "lin-cms"
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),

		kratos.Logger(logger),
		kratos.Server(hs))
}

func main() {
	os.RemoveAll("./log")
	conf := newConfig()
	logger := log.With(kzap.NewLogger(newLogger()),
		"service.id", id,
		"service.name", Name,

		"caller", log.DefaultCaller,
	)

	app, cleanup, err := NewWire(conf.Server, conf.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = app.Run(); err != nil {
		panic(err)
	}
}

func newLogger() *zap.Logger {
	f, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeSyncer := zapcore.AddSync(io.MultiWriter(f, os.Stdout))
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoder := zapcore.NewJSONEncoder(cfg)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	return zap.New(core)
}

func newConfig() *conf.Bootstrap {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return &bc
}
