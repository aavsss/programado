package schedule

import (
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
	scheduleRoute.GET("/", sc.ViewSchedule)
	scheduleRoute.POST("/add", sc.AddToSchedule)
	scheduleRoute.POST("/remove", sc.RemoveFromSchedule)
	scheduleRoute.PUT("/update", sc.UpdateSchedule)
}

func (sc *ScheduleController) ViewSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "To be connected"})
}

func (sc *ScheduleController) AddToSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "To be connected"})
}

func (sc *ScheduleController) RemoveFromSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "To be connected"})
}

func (sc *ScheduleController) UpdateSchedule(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "To be connected"})
}
