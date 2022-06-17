package request

import (
	"entry-task-web/pkg/code"
	"entry-task-web/pkg/log"
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GenRequestID 生成全局请求id
func GenRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID, err := nanoid.New()
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
