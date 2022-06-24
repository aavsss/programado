package schedule

import (
	"github.com/aavsss/programado/api/v1/schedule/models"
	"github.com/aavsss/programado/api/v1/schedule/repository"
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
	returnChannel := make(chan []repository.PeriodInDB)

	go sc.ScheduleService.ViewSchedule(returnChannel)
	sortedSchedule := <-returnChannel
	c.IndentedJSON(http.StatusOK, sortedSchedule)
}

func (sc *ScheduleController) addToSchedule(c *gin.Context) {
	returnChannel := make(chan bool)
	var period models.Period

	if err := c.BindJSON(&period); err != nil {
		return
	}

	go sc.ScheduleService.AddToSchedule(period, returnChannel)
	isSuccess := <-returnChannel

	c.IndentedJSON(
		http.StatusOK, isSuccess,
	)
}

func (sc *ScheduleController) removeFromSchedule(c *gin.Context) {
	returnChannel := make(chan bool)
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	go sc.ScheduleService.RemoveFromSchedule(id, returnChannel)
	isSuccess := <-returnChannel

	c.IndentedJSON(http.StatusOK, isSuccess)
}

func (sc *ScheduleController) updateSchedule(c *gin.Context) {
	returnChannel := make(chan bool)
	var period models.Period

	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	if err := c.BindJSON(&period); err != nil {
		return
	}

	go sc.ScheduleService.UpdateSchedule(id, period, returnChannel)
	isSuccess := <-returnChannel

	c.IndentedJSON(http.StatusOK, isSuccess)
}
