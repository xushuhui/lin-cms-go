package router

import (
	"context"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
	port           string
)


func httpListen() {
	if err := HttpSrvHandler.ListenAndServe(); err != nil {
		log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
	}
}
func httpsListen() {
	if err := HttpSrvHandler.ListenAndServeTLS("storage/cert.pem", "storage/key.pem"); err != nil {
		log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
	}
}
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
