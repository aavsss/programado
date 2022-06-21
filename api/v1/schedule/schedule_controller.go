package schedule

import (
	"github.com/aavsss/programado/api/v1/schedule/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

type ScheduleController struct {
	ScheduleService ScheduleService
}

func NewController(scheduleService ScheduleService) ScheduleController {
	return ScheduleController{
		ScheduleService: scheduleService,
	}
}

func (sc *ScheduleController) RegisterScheduleRoutes(rg *gin.RouterGroup) {
	scheduleRoute := rg.Group("/schedule")
	scheduleRoute.GET("/", sc.viewSchedule)
	scheduleRoute.POST("/add", sc.addToSchedule)
	scheduleRoute.POST("/remove", sc.removeFromSchedule)
	scheduleRoute.PUT("/update", sc.updateSchedule)
}

func (sc *ScheduleController) viewSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sc.ScheduleService.ViewSchedule())
}

func (sc *ScheduleController) addToSchedule(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK, sc.ScheduleService.AddToSchedule(
			models.Period{
				StartTime:   15,
				EndTime:     20,
				Description: "New Meeting",
			},
		),
	)
}

func (sc *ScheduleController) removeFromSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sc.ScheduleService.RemoveFromSchedule(2))
}

func (sc *ScheduleController) updateSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sc.ScheduleService.UpdateSchedule(3, models.Period{
		StartTime:   21,
		EndTime:     25,
		Description: "Evergreen in Las Collinas",
	}))
}
