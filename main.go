package main

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/routes"
	"time"

	_ "RaiJaiAPI_Golang/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    r := gin.Default()
    r.RedirectTrailingSlash = false

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

     r.Use(func(c *gin.Context) {
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
            return
        }
        c.Next()
    })
    
    database.ConnectDB()
    routes.SetupRoutes(r)
    
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}
