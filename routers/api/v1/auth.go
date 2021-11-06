package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"go-gin-example/models"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/merror"
	"go-gin-example/pkg/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := merror.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username, password)
			if err != nil {
				code = merror.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				data["token"] = token
				code = merror.SUCCESS
			}
		} else {
			code = merror.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": data,
	})
}
