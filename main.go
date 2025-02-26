package main

import (
	"to-do-list-app/config"
	"to-do-list-app/controller"
	"to-do-list-app/middleware"
	"to-do-list-app/repository"
	"to-do-list-app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//!! Database
	db := config.InitDB()

	//!! Repository
	userRepository := repository.NewUserRepository(db)
	checklistRepository := repository.NewChecklistRepository(db)

	//!! Service
	userService := service.NewUserService(userRepository)
	checklistService := service.NewChecklistService(checklistRepository)
	authService := service.NewAuthService()

	//!! Middleware
	authMiddleware := middleware.AuthMiddleware(authService, userService)

	//!! Controller
	userController := controller.NewUserController(userService, authService)
	checklistController := controller.NewChecklistController(checklistService)

	router := gin.Default()
	api := router.Group("/api/v1")

	//!! User Route
	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)

	//!! Checklist Route
	api.GET("/checklists", authMiddleware, checklistController.GetAllChecklists)
	api.GET("/checklists/:id", authMiddleware, checklistController.GetChecklist)
	api.POST("/checklists", authMiddleware, checklistController.CreateChecklist)
	api.DELETE("/checklists/:id", authMiddleware, checklistController.DeleteChecklist)

	router.Run()
}
