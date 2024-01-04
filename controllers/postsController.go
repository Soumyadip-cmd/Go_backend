package controllers

import (
	// "strconv"

	"github.com/CVWO/go-crud/database"
	"github.com/CVWO/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Content: body.Content}

	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post

	database.DB.Order("ID desc").Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id of the url
	id := c.Param("id")

	// Get the post
	var post models.Post
	database.DB.First(&post, id)

	// Respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off the url
	id := c.Param("id")

	// Get data off req
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	//Find post to update
	var post models.Post
	database.DB.First(&post, id)

	//Update post
	database.DB.Model(&post).Updates(models.Post{
		Title:   body.Title,
		Content: body.Content,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the post
	database.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
