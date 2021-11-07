package app

import (
	"github.com/gin-gonic/gin"

	"github.com/L-Chao/go-gin-example/pkg/merror"
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
