package master

import (
	"net/http"

	"github.com/aavsss/programado/api/v1/schedule/repository"
	"github.com/gin-gonic/gin"
)

type MasterController struct {
	MasterService MasterService
}

func NewController(masterService MasterService) MasterController {
	return MasterController{
		MasterService: masterService,
	}
}

func (mc *MasterController) RegisterMasterRoutes(rg *gin.RouterGroup) {
	masterRoute := rg.Group("/master")
	masterRoute.GET("/requests", mc.viewRequests)
	masterRoute.POST("/review_requests", mc.reviewRequests)
}

func (mc *MasterController) viewRequests(c *gin.Context) {
	masterId, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Master id is missing"})
		return
	}

	returnChannel := make(chan []repository.RequestInDB)
	go mc.MasterService.viewRequests(masterId, returnChannel)
	requests := <-returnChannel

	c.IndentedJSON(http.StatusOK, requests)
}

func (mc *MasterController) reviewRequests(c *gin.Context) {
	scheduleId, ok := c.GetQuery("schedule_id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Master id is missing"})
		return
	}

	go mc.MasterService.reviewRequest("2", scheduleId, true)
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Not implemented correctly"})

}
