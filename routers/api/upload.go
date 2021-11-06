package api

import (
	"net/http"

	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/merror"
	"go-gin-example/pkg/upload"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	code := merror.SUCCESS
	data := make(map[string]interface{})

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = merror.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  merror.GetMsg(code),
			"data": data,
		})
	}
	if image == nil {
		code = merror.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = merror.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = merror.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = merror.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": data,
	})
}
