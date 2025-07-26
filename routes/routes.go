package routes

import (
        "RaiJaiAPI_Golang/controllers"
        "RaiJaiAPI_Golang/middleware"

        "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    public := r.Group("/api")
    {
        public.POST("/login", controllers.Login)
        public.POST("/users", controllers.CreateUser)
    }

    auth := r.Group("/api")
    auth.Use(middleware.AuthMiddleware())

    user := auth.Group("/users")
    {
        user.GET("/", controllers.GetUsers)
        user.PUT("/:id", controllers.UpdateUser)
        user.DELETE("/:id", controllers.DeleteUser)
        user.GET("/:id", controllers.GetUserByID)
    }

    typeGroup := auth.Group("/types")
    {
        typeGroup.GET("/", controllers.GetTypes)
        typeGroup.POST("/", controllers.CreateType)
        typeGroup.PUT("/:id", controllers.UpdateType)
        typeGroup.DELETE("/:id", controllers.DeleteType)
        typeGroup.GET("/:id", controllers.GetTypeByID)
    }

    category := auth.Group("/categories")
    {
        category.GET("/", controllers.GetCategories)
        category.POST("/", controllers.CreateCategory)
        category.PUT("/:id", controllers.UpdateCategory)
        category.DELETE("/:id", controllers.DeleteCategory)
        category.GET("/:id", controllers.GetCategory)
    }

    transaction := auth.Group("/transactions")
    {
        transaction.GET("/", controllers.GetTransactions)
        transaction.POST("/", controllers.CreateTransaction)
        transaction.PUT("/:id", controllers.UpdateTransaction)
        transaction.DELETE("/:id", controllers.DeleteTransaction)
        transaction.GET("/:id", controllers.GetTransaction)
    }
}
