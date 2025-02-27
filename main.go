package main

import (
	"github.com/ipoluianov/cetuspools/app"
	"github.com/ipoluianov/cetuspools/application"
	"github.com/ipoluianov/cetuspools/logger"
)

func main() {
	application.Name = "cetuspools"
	application.ServiceName = "cetuspools"
	application.ServiceDisplayName = "cetuspools"
	application.ServiceDescription = "cetuspools"
	application.ServiceRunFunc = app.RunAsService
	application.ServiceStopFunc = app.StopService

	logger.Init(logger.CurrentExePath() + "/logs")

	if !application.TryService() {
		app.RunDesktop()
	}
}
