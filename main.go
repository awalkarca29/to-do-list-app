package main

import (
	"to-do-list-app/config"
	"to-do-list-app/controller"
	"to-do-list-app/repository"
	"to-do-list-app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//!! Database
	db := config.InitDB()

	//!! Repository
	userRepository := repository.NewUserRepository(db)

	//!! Service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()

	//!! Middleware
	// authMiddleware := middleware.AuthMiddleware(authService, userService)

	//!! Controller
	userController := controller.NewUserController(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	//!! User Route
	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)

	//!! Event Route
	// api.GET("/events", eventController.GetAllEvents)
	// api.GET("/events/:id", eventController.GetEvent)
	// api.POST("/events", authMiddlewareAdmin, eventController.CreateEvent)
	// api.PUT("/events/:id", authMiddlewareAdmin, eventController.UpdateEvent)
	// api.DELETE("/events/:id", authMiddlewareAdmin, eventController.DeleteEvent)

	router.Run()
}
