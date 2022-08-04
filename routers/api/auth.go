package api

import (
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/app"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/code"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/rpc"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/util"

	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth 登陆认证
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	id, ok := c.Get("requestId")
	if !ok {
		log.Errorf("get requestId failed")
		appG.Response(http.StatusInternalServerError, code.Error, nil)
		return
	}
	requestID := id.(string)
	valid := validation.Validation{}

	name := c.PostForm("name")
	password := c.PostForm("password")

	a := auth{Username: name, Password: password}
	ok, _ = valid.Valid(&a)

	if !ok {
		app.MarkErrors(requestID, valid.Errors)
		appG.Response(http.StatusBadRequest, code.InvalidParams, nil)
		return
	}

	nickname, image, err := rpc.Auth(requestID, name, password)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, nil)
		return
	}

	token, err := util.GenerateToken(name, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ErrorAuthToken, nil)
		return
	}

	appG.Response(http.StatusOK, code.Success, map[string]string{
		"token":    token,
		"nickname": nickname,
		"image":    image,
	})
}
