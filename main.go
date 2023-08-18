package main

import (
	"todo-app/database"
	"todo-app/docs"
	"todo-app/model"
	"todo-app/router"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func main() {
	d := database.Connect()
	d.AutoMigrate(&model.Todo{})
	r := gin.Default()
	router.LoadAuthRouter(r)
	router.LoadTodoRouter(r)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
