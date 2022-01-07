package app

import (
	"microservices/controllers"
)

func mapUrls() {
	router.POST("/microservices", controllers.StudentsController.Create)
	router.GET("/microservices/:id", controllers.StudentsController.Get)
}
