package routers

import (
	"net/http"

	"github.com/Manikandan-Parasuraman/authentication-system-go/src/controllers"
	"github.com/Manikandan-Parasuraman/authentication-system-go/src/models"

	"github.com/gin-gonic/gin"
)

func UserRouter(context *gin.Context) {
    var inputData models.CreateUserRequest
    if err := context.ShouldBindJSON(&inputData); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Please check input and request again."})
        return
	}
	controllers.UserCtrl(inputData, context)
}