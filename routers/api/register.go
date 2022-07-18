package api

import (
	"entry-task-web/pkg/app"
	"entry-task-web/pkg/code"
	"entry-task-web/pkg/log"
	"entry-task-web/pkg/rpc"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerInfo struct {
	Name     string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MinSize(6); MaxSize(50)"`
}

// Register 注册
func Register(c *gin.Context) {
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

	a := registerInfo{Name: name, Password: password}
	ok, _ = valid.Valid(&a)
	if !ok {
		app.MarkErrors(requestID, valid.Errors)
		appG.Response(http.StatusBadRequest, code.InvalidParams, nil)
		return
	}

	err := rpc.Register(requestID, name, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ErrorRegisterFailed, nil)
		return
	}

	appG.Response(http.StatusOK, code.Success, map[string]string{})
}