package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/L-Chao/go-gin-example/middleware/jwt"
	"github.com/L-Chao/go-gin-example/pkg/setting"
	"github.com/L-Chao/go-gin-example/pkg/upload"
	"github.com/L-Chao/go-gin-example/routers/api"
	v1 "github.com/L-Chao/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/auth", v1.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTag)

		apiv1.PUT("/tags/:id", v1.EditTag)

		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)

		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.POST("/articles", v1.AddArticle)

		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
