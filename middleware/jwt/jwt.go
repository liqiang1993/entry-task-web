package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/code"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/util"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resCode code.ErrorStruct
		var data interface{}
		resCode = code.Success

		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		if token == "" {
			resCode = code.InvalidParams
		} else {
			raw, err := util.ParseToken(token)
			if err == nil {
				c.Set("name", raw.Username)
			} else {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					resCode = code.ErrorAuthCheckTokenTimeout
				default:
					resCode = code.ErrorAuthCheckTokenFail
				}
			}
		}

		if resCode != code.Success {
			var requestID string
			id, ok := c.Get("requestId")
			if !ok {
				log.Errorf("get requestId failed")
			} else {
				requestID = id.(string)
			}

			log.Warnf("%s auth failed:%+v", requestID, resCode)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": resCode.Code,
				"msg":  resCode.Msg,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
