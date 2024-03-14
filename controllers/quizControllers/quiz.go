package quizControllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/web-service-gin/models"
)

// GetAllquizs retrieves all quizs from the MongoDB collection
func GetAllquizes(c *gin.Context) {
	collection := c.MustGet("quizCollection").(*mongo.Collection)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var quizes []models.Quiz
	if err := cursor.All(context.Background(), &quizes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizes)
}
