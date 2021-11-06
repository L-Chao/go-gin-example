package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"go-gin-example/models"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/merror"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/utils"
)

func GetArticle(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := merror.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(int(id)) {
			data = models.GetArticle(int(id))
			code = merror.SUCCESS
		} else {
			code = merror.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int64 = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.ParseInt(arg, 10, 32)
		maps["state"] = int(state)

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int64 = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId, _ = strconv.ParseInt(arg, 10, 32)
		maps["tag_id"] = int(tagId)

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := merror.INVALID_PARAMS
	if !valid.HasErrors() {
		code = merror.SUCCESS

		data["lists"] = models.GetArticles(utils.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)

	} else {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	tagId, _ := strconv.ParseInt(c.Query("tag_id"), 10, 32)
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state, _ := strconv.ParseInt(c.DefaultQuery("state", "0"), 10, 32)

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := merror.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByID(int(tagId)) {
			data := make(map[string]interface{})
			data["tag_id"] = int(tagId)
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = int(state)

			models.AddArticle(data)
			code = merror.SUCCESS
		} else {
			code = merror.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id, _ := strconv.ParseInt((c.Param("id")), 10, 32)
	tagId, _ := strconv.ParseInt(c.Query("tag_id"), 10, 32)
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int64 = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.ParseInt(arg, 10, 32)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := merror.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(int(id)) {
			if models.ExistTagByID(int(tagId)) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = int(tagId)
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(int(id), data)
				code = merror.SUCCESS
			} else {
				code = merror.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = merror.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := merror.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(int(id)) {
			models.DeleteArticle(int(id))
			code = merror.SUCCESS
		} else {
			code = merror.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": make(map[string]string),
	})
}
