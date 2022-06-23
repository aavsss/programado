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
	var period models.Period

	if err := c.BindJSON(&period); err != nil {
		return
	}

	c.IndentedJSON(
		http.StatusOK, sc.ScheduleService.AddToSchedule(period),
	)
}

func (sc *ScheduleController) removeFromSchedule(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	c.IndentedJSON(http.StatusOK, sc.ScheduleService.RemoveFromSchedule(id))
}

func (sc *ScheduleController) updateSchedule(c *gin.Context) {
	var period models.Period

	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	if err := c.BindJSON(&period); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, sc.ScheduleService.UpdateSchedule(id, period))
}
