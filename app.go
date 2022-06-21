package main

import (
	"github.com/aavsss/programado/api/v1/schedule"
	"github.com/aavsss/programado/api/v1/version"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	scheduleController schedule.ScheduleController
	scheduleService    schedule.ScheduleService

	versionController version.VersionController
	versionService    version.VersionService
)

func init() {
	server = gin.Default()

	scheduleService = schedule.NewService()
	scheduleController = schedule.NewController(scheduleService)

	versionService = version.NewService()
	versionController = version.NewController(versionService)
}
