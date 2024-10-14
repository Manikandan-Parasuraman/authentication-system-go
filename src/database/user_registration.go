package database

import (
	"context"
	"net/http"
	"time"

	"github.com/Manikandan-Parasuraman/authentication-system-go/src/models"
	"github.com/Manikandan-Parasuraman/authentication-system-go/src/utility"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user.
func UserDatabase(data models.User, c *gin.Context)(*models.User, error){
	
    // Get and validate the mongodb connections
    client , err := utility.ConnectToMongoDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return nil, err
    }

    // Create the collections
    collection := client.Collection("user_management")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Validate the input data with database
	err = collection.FindOne(context.Background(), bson.M{"email": data.Email}).Decode(&data)
	if err == nil {
		return nil, err
	} else {
        // Data insertion
        result, err := collection.InsertOne(ctx, data)
        if err != nil {
            // Get the inserted ID
            insertedID := result.InsertedID
            
            // Read inserted data
            err := collection.FindOne(context.Background(), bson.M{"_id": insertedID}).Decode(&data)
            if err != nil {
                return nil, err
            }
        }
    }
    return &data, err
}
