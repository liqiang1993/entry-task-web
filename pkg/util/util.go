package util

import (
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/setting"
)

// InitUtil 初始化工具
func InitUtil() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
