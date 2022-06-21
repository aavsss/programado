package main

func main() {
	basePath := server.Group("v1")
	scheduleController.RegisterScheduleRoutes(basePath)
	versionController.RegisterVersionRoutes(basePath)

	server.Run("localhost:8080")
}
