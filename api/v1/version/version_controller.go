package version

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type VersionController struct {
	VersionService VersionService
}

func NewController(versionService VersionService) VersionController {
	return VersionController{
		VersionService: versionService,
	}
}

func (vc *VersionController) RegisterVersionRoutes(rg *gin.RouterGroup) {
	versionRoute := rg.Group("version")
	versionRoute.GET("/", vc.GetAppVersion)
}

func (vc *VersionController) GetAppVersion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vc.VersionService.GetAppVersion())
}
