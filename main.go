package main

func main() {
	basePath := server.Group("v1")
	clientController.RegisterUserRoutes(basePath)
	scheduleController.RegisterScheduleRoutes(basePath)
	versionController.RegisterVersionRoutes(basePath)
	masterController.RegisterMasterRoutes(basePath)

	server.Run("localhost:8080")
}
