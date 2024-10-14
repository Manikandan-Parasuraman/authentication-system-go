package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInitiator(server *gin.Engine) {
	server.POST("/user", UserRouter)
	grouping := server.Group("/api/v1")
	grouping.GET("/", getEvent)
}


// PingExample godoc
// @Summary ping example
// @Schemes 
// @Description do ping
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/v1 [get]
func getEvent(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{"message":"Hello World..!"})
}