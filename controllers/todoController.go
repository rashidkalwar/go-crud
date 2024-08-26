package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rashidkalwar/go-crud/initializers"
	"github.com/rashidkalwar/go-crud/models"
)

func TodosCreate(c *gin.Context) {
	/*
		We are not focusing on validations for now!
	*/
	// Get data from req body
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	// Create a todo
	todo := models.Todo{Content: body.Content, Status: body.Status}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosIndex(c *gin.Context) {
	/*
		Not worrying about security and auth based retrieval
	*/

	// Get all the todos
	var todos []models.Todo
	initializers.DB.Find(&todos)

	// Return todos in response
	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func TodosShow(c *gin.Context) {
	/*
		Not worrying about security and auth based retrieval
	*/
	// Get id from URL param
	id := c.Param("id")

	// Get a sing todo
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Return todo in response
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosUpdate(c *gin.Context) {
	/*
		Not worrying about security and auth based updating
	*/
	// Get id from URL param
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	// Get a single todo that we want to update
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Update it
	initializers.DB.Model(&todo).Updates(models.Todo{
		Content: body.Content,
		Status:  body.Status,
	})

	// Return response
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosDelete(c *gin.Context) {
	/*
		Not worrying about security and auth based deletion
	*/
	// Get id from URL param
	id := c.Param("id")

	// Delete the Todo
	initializers.DB.Delete(&models.Todo{}, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Todo removed Successfully",
	})
}
