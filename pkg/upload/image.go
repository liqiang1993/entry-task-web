package upload

import (
	"entry-task-web/pkg/file"
	"entry-task-web/pkg/setting"
	"strings"
)

// CheckImageSize 检查图片大小是否合法
func CheckImageSize(size int) bool {
	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImageExt 检查扩展是否合法
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExt {
		if strings.EqualFold(allowExt, ext) {
			return true
		}
	}
	return false
}