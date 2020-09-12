package main

import (
	"lin-cms-go/internal/router"
	"lin-cms-go/pkg/core"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	core.StartModule()

	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()

}
