package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/L-Chao/go-gin-example/models"
	"github.com/L-Chao/go-gin-example/pkg/merror"
	"github.com/L-Chao/go-gin-example/pkg/setting"
	"github.com/L-Chao/go-gin-example/pkg/utils"
	"github.com/astaxie/beego/validation"
)

// @Summary 获取文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int64 = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.ParseInt(arg, 10, 32)
		maps["state"] = int(state)
	}
	code := merror.SUCCESS
	data["lists"] = models.GetTags(utils.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state, _ := strconv.ParseInt(c.DefaultQuery("state", "0"), 10, 32)
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100个字符")
	valid.Required(createdBy, "created_by").Message("创建问不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态为只允许0或1")

	code := merror.INVALID_PARAMS

	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = merror.SUCCESS
			models.AddTag(name, int(state), createdBy)
		} else {
			code = merror.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int64 = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.ParseInt(arg, 10, 32)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := merror.INVALID_PARAMS

	if !valid.HasErrors() {
		code = merror.SUCCESS
		if models.ExistTagByID(int(id)) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(int(id), data)
		} else {
			code = merror.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": make(map[string]string),
	})

}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := merror.INVALID_PARAMS

	if !valid.HasErrors() {
		code = merror.SUCCESS
		if models.ExistTagByID(int(id)) {
			models.DeleteTag(int(id))
		} else {
			code = merror.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  merror.GetMsg(code),
		"data": map[string]string{},
	})
}
