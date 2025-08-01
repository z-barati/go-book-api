package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery middleware handles panics and returns 500 error
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
				"message": err,
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
} 