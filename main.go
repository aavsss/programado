package main

func main() {
	basePath := server.Group("v1")
	scheduleController.RegisterScheduleRoutes(basePath)

	server.Run("localhost:8080")
}
