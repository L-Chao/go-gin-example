package app

import (
	"go-gin-example/pkg/merror"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  merror.GetMsg(errCode),
		"data": data,
	})
}
