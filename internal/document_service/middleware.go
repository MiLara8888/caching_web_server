package documentservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		if c.Request.Header["Content-Type"] == nil || len(c.Request.Header["Content-Type"]) == 0 {
			c.AbortWithStatus(204)
			return
		}

		content_type := c.Request.Header["Content-Type"]
		if content_type[0] != "application/json" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

