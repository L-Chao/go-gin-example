package main

import (
	"fmt"
	"log"
	"syscall"

	"go-gin-example/pkg/setting"
	"go-gin-example/routers"

	"go-gin-example/cron"

	"github.com/fvbock/endless"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeOut
	endless.DefaultWriteTimeOut = setting.WriteTimeOut
	endless.DefaultMaxHeaderBytes = 1 << 20
	endpoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endpoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	go cron.CronMain()
	if err != nil {
		log.Printf("server error: %v", err)
	}
}
