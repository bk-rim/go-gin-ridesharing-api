package main

import (
	"github.com/gin-gonic/gin"

	"api/allo-dakar/controller"
	"api/allo-dakar/middleware"
	"api/allo-dakar/utils"

	"api/allo-dakar/database"
)

func main() {
	utils.LoadEnv(".env")
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	database.Connect()
	database.MigrateDb()
}

func serverApplication() {
	router := gin.Default()

	publicRoutesAuth := router.Group("/api/v1/auth")
	publicRoutesAuth.POST("/login", controller.AuthenticateUser)
	publicRoutesAuth.POST("/register", controller.CreateUser)

	publicRoutes := router.Group("/api/v1")
	publicRoutes.GET("/travels", controller.GetTravels)
	publicRoutes.GET("/travels/start/:start/end/:end", controller.GetTravelsByStartAndEnd)
	publicRoutes.GET("/travels/start/:start/end/:end/date/:date", controller.GetTravelsByStartAndEndAndDate)

	privateRoutes := router.Group("/api/v1")
	privateRoutes.Use(middleware.JwtAuth())
	privateRoutes.POST("/travels", controller.CreateTravel)
	privateRoutes.GET("/travels/idDriver/:idDriver", controller.GetTravelsByIdDriver)
	privateRoutes.DELETE("/travels/idTravel/:idTravel", controller.DeleteTravel)

	privateRoutesAdmin := router.Group("/api/v1/admin")
	privateRoutesAdmin.Use(middleware.JwtAuth())
	privateRoutesAdmin.GET("/users", controller.GetUsers)

	router.Run("localhost:8080")
}
