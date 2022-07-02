package master

import (
	"net/http"

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
}

func (mc *MasterController) viewRequests(c *gin.Context) {
	masterId, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Master id is missing"})
		return
	}

	c.IndentedJSON(http.StatusOK, mc.MasterService.viewRequests(masterId))
}
