package auth

import (
	"fmt"
	"net/http"

	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

type MiddlewareOptions struct {
	IsAdmin bool
}

func (a Auth) Middleware(options *MiddlewareOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, err := parseToken(c, a.cfg.Jwt.AccessTokenPrivateKey)
		if err != nil {
			return
		}

		currentUser, err := a.userUC.FindByID(fmt.Sprint(sub.UserID))

		if options != nil {
			if options.IsAdmin && !currentUser.IsSuperUser {
				c.AbortWithStatusJSON(http.StatusForbidden, api.NewError("not enough rights"))
				return
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.NewError("the user belonging to this token no logger exists"))
			return
		}

		c.Set(keyUserStorageCtx, *currentUser)
		c.Next()
	}
}
