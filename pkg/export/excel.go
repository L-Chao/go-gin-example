package export

import (
	"github.com/L-Chao/go-gin-example/pkg/setting"
)

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.AppSetting.ExportSavePath + GetExcelPath()
}
