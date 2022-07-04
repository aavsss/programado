package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	ClientService ClientService
}

func NewController(clientService ClientService) ClientController {
	return ClientController{
		ClientService: clientService,
	}
}

func (cc *ClientController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/client")
	userRoute.POST("/add_meeting", cc.requestNewMeeting)
}

func (cc *ClientController) requestNewMeeting(c *gin.Context) {
	returnChannel := make(chan bool)
	go cc.ClientService.RequestNewMeeting(returnChannel)
	isSuccess := <-returnChannel
	c.IndentedJSON(http.StatusOK, isSuccess)
}
