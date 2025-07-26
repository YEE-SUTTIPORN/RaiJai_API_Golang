package middleware

import (
	"net/http"
	"strings"

	"RaiJaiAPI_Golang/models"
	"RaiJaiAPI_Golang/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.JsonResponse{Success: false, Message: "Unauthorized"})
			return
		}
		token := strings.TrimPrefix(h, "Bearer ")
		uid, err := utils.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.JsonResponse{Success: false, Message: "Unauthorized"})
			return
		}
		c.Set("userID", uid)
		c.Next()
	}
}
