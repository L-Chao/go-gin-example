package v1

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/L-Chao/go-gin-example/pkg/app"
	"github.com/L-Chao/go-gin-example/pkg/merror"
	"github.com/L-Chao/go-gin-example/pkg/setting"
	"github.com/L-Chao/go-gin-example/pkg/utils"
	"github.com/L-Chao/go-gin-example/service/article_service"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, merror.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: int(id)}
	exist, err := articleService.ExistArticleByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, merror.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, merror.SUCCESS, article)
}

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var state int64 = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.ParseInt(arg, 10, 32)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int64 = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId, _ = strconv.ParseInt(arg, 10, 32)
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	articleService := article_service.Article{
		TagID:    int(tagId),
		State:    int(state),
		PageNum:  utils.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}
	artilces, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_GET_ALL_ARTICLE_FAIL, nil)
		return
	}
	data := map[string]interface{}{
		"list":  artilces,
		"total": total,
	}
	appG.Response(http.StatusOK, merror.SUCCESS, data)
}

func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	tagId, _ := strconv.ParseInt(c.Query("tag_id"), 10, 32)
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state, _ := strconv.ParseInt(c.DefaultQuery("state", "0"), 10, 32)
	coverImageUrl := c.Query("cover_image_url")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(coverImageUrl, "cover_image_url").Message("封面图片URL不能为空")
	valid.MaxSize(coverImageUrl, 1000, "cover_image_url")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, merror.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		TagID:         int(tagId),
		Title:         title,
		Desc:          desc,
		Content:       content,
		CreatedBy:     createdBy,
		State:         int(state),
		CoverImageUrl: coverImageUrl,
	}
	err := articleService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, merror.SUCCESS, nil)
}

//修改文章
func EditArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id, _ := strconv.ParseInt((c.Param("id")), 10, 32)
	tagId, _ := strconv.ParseInt(c.Query("tag_id"), 10, 32)
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	coverImageUrl := c.Query("cover_image_url")

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
	// valid.Required(coverImageUrl, "cover_image_url").Message("封面图片URL不能为空")
	valid.MaxSize(coverImageUrl, 1000, "cover_image_url")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, merror.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		ID:            int(id),
		TagID:         int(tagId),
		Title:         title,
		Desc:          desc,
		Content:       content,
		ModifiedBy:    modifiedBy,
		State:         int(state),
		CoverImageUrl: coverImageUrl,
	}
	err := articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_EDIE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, merror.SUCCESS, nil)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, merror.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		ID: int(id),
	}
	err := articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, merror.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, merror.SUCCESS, nil)
}
