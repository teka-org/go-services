package avatarControllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/web-service-gin/models"
)

// GetAllAvatars retrieves all avatars from the MongoDB collection
func GetAllAvatars(c *gin.Context) {
	collection := c.MustGet("avatarCollection").(*mongo.Collection)
	// collection := c.MustGet("avatar").(*mongo.Collection)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var avatars []models.Avatar
	if err := cursor.All(context.Background(), &avatars); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, avatars)
}

// GetADetailAvatars retrieves detail avatars from the MongoDB collection
// func GetDetailAvatar(c *gin.Context, error) {
// 	collection := c.MustGet("collection").(*mongo.Collection)

// 	id := c.Param("id")

// 	var result bson.M
// 	err = collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)
// 	if err == mongo.ErrNoDocuments {
// 		fmt.Printf("No avatar was found")
// 		return
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	avatar, err := json.MarshalIndent(result, "", "    ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", avatar)
// }

// func GetAvatarDetail(c *gin.Context) {
//     collection := c.MustGet("collection").(*mongo.Collection)

//     // Get the avatar ID from the request parameters
//     avatarID := c.Param("id")

//     // Define a filter to find the avatar by its ID
//     filter := bson.M{"_id": bson.ObjectIdHex(avatarID)}

//     // Find the avatar by ID
//     var avatar models.Avatar
//     err := collection.FindOne(context.Background(), filter).Decode(&avatar)
//     if err != nil {
//         if err == mongo.ErrNoDocuments {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
//             return
//         }
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Return the avatar detail
//     c.JSON(http.StatusOK, avatar)
// }
