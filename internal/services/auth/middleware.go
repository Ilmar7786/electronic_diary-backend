package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newMiddleware(privateKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, err := parseToken(c, privateKey)
		if err != nil {
			return
		}

		userId := fmt.Sprint(sub.UserID)
		if userId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "the user belonging to this token no logger exists"})
			return
		}

		c.Set(keyUserStorageCtx, userId)
		c.Next()
	}
}

func getMustGetAuth(ctx *gin.Context) string {
	return ctx.MustGet(keyUserStorageCtx).(string)
}
