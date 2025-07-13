package main

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/routes"

	_ "RaiJaiAPI_Golang/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    r := gin.Default()

    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    database.ConnectDB()
    routes.SetupRoutes(r)

    r.Run(":8080")
}
