package main

import (
	"github.com/frederico-neres/my-own-log-management--litelogs/adapter/input/udp"
	"github.com/frederico-neres/my-own-log-management--litelogs/adapter/output/repository"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/service"
)

func main() {
	cloverDbRepository := repository.NewCloverDbRepository()
	saveLogService := service.NewSaveLogService(cloverDbRepository)
	server := udp.NewUdpServer(saveLogService)
	server.Listen()
}
