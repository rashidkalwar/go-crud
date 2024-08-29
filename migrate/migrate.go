package main

import (
	"github.com/rashidkalwar/go-crud/initializers"
	"github.com/rashidkalwar/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
