package main

import (
	"github.com/aavsss/programado/api/v1/client"
	"github.com/aavsss/programado/api/v1/master"
	"github.com/aavsss/programado/api/v1/schedule"
	"github.com/aavsss/programado/api/v1/version"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	clientController client.ClientController
	clientService    client.ClientService

	scheduleController schedule.ScheduleController
	scheduleService    schedule.ScheduleService

	versionController version.VersionController
	versionService    version.VersionService

	masterController master.MasterController
	masterService    master.MasterService
)

func init() {
	server = gin.Default()

	scheduleService = schedule.NewService()
	scheduleController = schedule.NewController(scheduleService)

	versionService = version.NewService()
	versionController = version.NewController(versionService)

	clientService = client.NewService(scheduleService)
	clientController = client.NewController(clientService)

	masterService = master.NewService()
	masterController = master.NewController(masterService)

}
