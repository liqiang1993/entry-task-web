package request

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/code"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/log"
	nanoid "github.com/matoous/go-nanoid"
	"net/http"
)

// GenRequestID 生成全局请求id
func GenRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID, err := nanoid.Nanoid()
		if err != nil {
			log.Warnf("gen nanoid failed:%s", err)

			var data interface{}
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": code.Error.Code,
				"msg":  code.Error.Msg,
				"data": data,
			})

			c.Abort()
			return
		}
		c.Set("requestId", requestID)
		c.Next()
	}
}
