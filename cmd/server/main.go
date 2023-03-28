package main

import (
	"love-scroll-api/internal/config"
	"love-scroll-api/internal/handler"
	"love-scroll-api/internal/middleware"
	"love-scroll-api/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Middleware
	r.Use(middleware.CORS())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	apiRouter := r.Group("/api/v1")

	// User routes
	apiRouter.POST("/register", handler.RegisterUser)
	apiRouter.POST("/login", handler.LoginUser)

	authorized := apiRouter.Group("/protected")
	authorized.Use(middleware.JWTAuth(cfg.JWT.Secret))
	{
		// Add your authorized routes here
		authorized.GET("/users/:username", handler.GetUserHandler)
		authorized.PUT("/users/:id", handler.UpdateUserHandler)
		authorized.DELETE("/users/:id", handler.DeleteUserHandler)
		authorized.GET("/users", handler.ListUsersHandler)
	}




	err = r.Run(cfg.Server.Address)
	if err != nil {
		panic(err)
	}
}
