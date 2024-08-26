package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rashidkalwar/go-crud/initializers"
	"github.com/rashidkalwar/go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	r := gin.Default()

	// Todo Routes
	routes.TodoRoutes(r)

	r.Run()
}
