package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-gin-example/pkg/merror"
	"go-gin-example/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = merror.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = merror.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = merror.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = merror.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != merror.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  merror.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
