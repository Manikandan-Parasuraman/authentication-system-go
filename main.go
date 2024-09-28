package main

import (
	"log"

	"github.com/Manikandan-Parasuraman/authentication-system-go/docs"
	"github.com/Manikandan-Parasuraman/authentication-system-go/src/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authentication System
// @version 1.0
// @description Authentication system api documentation
// @contact.name API Support
// @contact.email tech.cod.mk13@gmail.com
// @license.name GNU General Public License v3.0
// @license.url https://www.gnu.org/licenses/gpl.htm
// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// router server initiate
	server := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routers.RouterInitiator(server)
	server.Run(":8080")
}
