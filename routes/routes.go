package routes

import (
	"RaiJaiAPI_Golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    user := r.Group("api/users")
    {
        user.GET("/", controllers.GetUsers)
        user.POST("/", controllers.CreateUser)
        user.PUT("/:id", controllers.UpdateUser)
        user.DELETE("/:id", controllers.DeleteUser)
        user.GET("/:id", controllers.GetUserByID)
    }

    typeGroup := r.Group("api/types")
    {
        typeGroup.GET("/", controllers.GetTypes)
        typeGroup.POST("/", controllers.CreateType)
        typeGroup.PUT("/:id", controllers.UpdateType)
        typeGroup.DELETE("/:id", controllers.DeleteType)
        typeGroup.GET("/:id", controllers.GetTypeByID)
    }

    category := r.Group("api/categories")
    {
        category.GET("/", controllers.GetCategories)
        category.POST("/", controllers.CreateCategory)
        category.PUT("/:id", controllers.UpdateCategory)
        category.DELETE("/:id", controllers.DeleteCategory)
        category.GET("/:id", controllers.GetCategory)
    }

    transaction := r.Group("api/transactions")
    {
        transaction.GET("/", controllers.GetTransactions)
        transaction.POST("/", controllers.CreateTransaction)
        transaction.PUT("/:id", controllers.UpdateTransaction)
        transaction.DELETE("/:id", controllers.DeleteTransaction)
        transaction.GET("/:id", controllers.GetTransaction)
    }
}
