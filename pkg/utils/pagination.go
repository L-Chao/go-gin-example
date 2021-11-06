package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-gin-example/pkg/setting"
)

func GetPage(c *gin.Context) int {
	var result int = 0
	pageStr := c.Query("page")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	if page > 0 {
		result = (int(page) - 1) * setting.AppSetting.PageSize
	}
	return result
}
