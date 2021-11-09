package controllers

import (
	"beego_start/models"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/globalsign/mgo/bson"
)

// DemoController 某对象相关操作
type DemoController struct {
	beego.Controller
}

// Post 创建某对象
// @Description 创建一个新某对象文档
// @Param	body	body	models.DemoEditRequest	true	"body 内容"
// @Success	200	{string}	models.Demo.Id
// @router / [post]
func (c *DemoController) Post() {
	var demo models.Demo
	json.Unmarshal(c.Ctx.Input.RequestBody, &demo)
	valid := validation.Validation{}
	valid.Required(demo.Score, "得分").Message("不能为空")
	valid.Range(demo.Score, 0, 100, "得分").Message("需在0~100之间")
	valid.Required(demo.PlayerName, "选手姓名").Message("不能为空")
	if ValidationInspection(c.Ctx, valid) != nil {
		return
	}
	id, err := models.AddDemo(demo)
	if err != nil {
		logs.Error(err.Error())
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = map[string]string{"id": id}
	c.ServeJSON()
}

// Get 获取某对象
// @Description 根据ID获取某对象文档
// @Param	id	path	string	true	"某对象ID"
// @Success 200 {object} models.DemoEditRequest
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router /:id [get]
func (c *DemoController) Get() {
	id := c.GetString(":id")
	if bson.IsObjectIdHex(id) == false {
		c.CustomAbort(400, "某对象ID的格式不正确")
	}
	demo, err := models.GetDemo(id)
	if err != nil {
		logs.Error(err.Error())
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = demo
	c.ServeJSON()
}

// GetAll 获取全部某对象
// @Description 获取全部某对象文档
// @Success 200 {object} models.DemoEditRequest
// @Failure 500 服务异常提示字符串
// @router / [get]
func (c *DemoController) GetAll() {
	demo, err := models.GetAllDemo()
	if err != nil {
		logs.Error(err.Error())
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = demo
	c.ServeJSON()
}

// Put 更新某对象
// @Description 根据ID更新某对象文档
// @Param	id	path	string	true	"某对象ID"
// @Param	body	body	models.DemoDocuUpdate	true	"body content"
// @Success 200 {string} Update Success!
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router /:id [put]
func (c *DemoController) Put() {
	id := c.GetString(":id")
	if bson.IsObjectIdHex(id) == false {
		c.CustomAbort(400, "某对象ID的格式不正确")
	}
	var object models.DemoDocuUpdate
	json.Unmarshal(c.Ctx.Input.RequestBody, &object)
	valid := validation.Validation{}
	valid.Required(object.Score, "分数").Message("不能为空")
	valid.Range(object.Score, 0, 100, "分数").Message("需在0~100之间")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.CustomAbort(400, fmt.Sprintf("%s%s", err.Key, err.Message))
		}
	}
	err := models.UpdateDemo(id, object)
	if err != nil {
		logs.Error(err.Error())
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = "Update Success!"
	c.ServeJSON()
}

// Delete 删除某对象
// @Description 根据ID删除某对象文档
// @Param	id	path	string	true	"某对象ID"
// @Success 200 {string} Delete Success!
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router /:id [delete]
func (c *DemoController) Delete() {
	id := c.GetString(":id")
	if bson.IsObjectIdHex(id) == false {
		c.CustomAbort(400, "某对象ID的格式不正确")
	}
	err := models.DeleteDemo(id)
	if err != nil {
		logs.Error(err.Error())
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = "Delete Success!"
	c.ServeJSON()
}
