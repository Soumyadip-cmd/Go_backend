package main

import (
	"github.com/CVWO/go-crud/controllers"
	"github.com/CVWO/go-crud/database"
	"github.com/CVWO/go-crud/initializers"
	"github.com/CVWO/go-crud/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.Connect()
	database.ConnectToDb()
}

// func optionsHandler(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
// 	c.Header("Access-Control-Allow-Credentials", "true")
// 	c.AbortWithStatus(200)
// }

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"}
	config.AllowCredentials = true
	// r.OPTIONS("/someGetPath", optionsHandler)
	// r.OPTIONS("/somePostPath", optionsHandler)
	// r.OPTIONS("/somePutPath", optionsHandler)
	// r.OPTIONS("/someDeletePath", optionsHandler)
	r.Use(cors.New(config))
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PATCH("/posts/:id", middleware.RequireAuth, controllers.PostsUpdate)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.PostsDelete)
	// Routes for authentication
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	//r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080
}
