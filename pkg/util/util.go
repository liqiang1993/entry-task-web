package util

import (
	"entry-task-web/pkg/setting"
)

// InitUtil 初始化工具
func InitUtil() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
