package server

import (
	"encoding/json"
	"time"

	"lin-cms-go/api"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"google.golang.org/grpc/encoding"
)

// MiddlewareCors 设置跨域请求头

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, s *service.CmsService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(

			recovery.Recovery(),

			logging.Server(logger),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "content-type", "Content-Type", "Authorization", "tag"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowCredentials(),
			handlers.AllowedOrigins([]string{"*"}),
		)),
		// http.ResponseEncoder(DefaultResponseEncoder),
	}

	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(5*time.Second))
	}

	srv := http.NewServer(opts...)

	api.RegisterCmsHTTPServer(srv, s)
	return srv
}

type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func DefaultResponseEncoder(w http.ResponseWriter, r *http.Request, v any) error {
	if v == nil {
		return nil
	}

	data, err := encoding.GetCodec("json").Marshal(v)
	if err != nil {
		return err
	}

	bs, _ := json.Marshal(&Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	if err != nil {
		return err
	}
	return nil
}
