package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Logging the request", time.Now())
		ctx := context.WithValue(c.Request.Context(), "userIds", "123")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
