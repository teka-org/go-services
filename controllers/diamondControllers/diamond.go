package diamondControllers

import (
	"context"
	"net/http"

	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAlldiamonds retrieves all diamonds from the MongoDB collection
func GetAllDiamonds(c *gin.Context) {
	collection := c.MustGet("diamondCollection").(*mongo.Collection)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var diamonds []models.Diamond
	if err := cursor.All(context.Background(), &diamonds); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diamonds)
}
