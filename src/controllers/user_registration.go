package controllers

import (
	"net/http"
	"time"

	"github.com/Manikandan-Parasuraman/authentication-system-go/src/database"
	"github.com/Manikandan-Parasuraman/authentication-system-go/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func UserCtrl(inputData models.CreateUserRequest, context *gin.Context) {  
    
    data := models.User{
        Name:      inputData.Name,
        Email:     inputData.Email,
        Password:  inputData.Password,
        CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
    }

    userData, err := database.UserDatabase(data, context)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    } else if userData == nil {
        context.JSON(http.StatusAlreadyReported, gin.H{"error": "Email already exists"})
        return
    }
    context.JSON(http.StatusCreated, models.UserResponse{
        ID:    userData.ID.Hex(),
        Name:  userData.Name,
        Email: userData.Email,
    })
}