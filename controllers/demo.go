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

// 某对象相关操作
type DemoController struct {
	beego.Controller
}

// @Title 创建某对象
// @Description 创建一个新某对象文档
// @Param	body	body	models.DemoDocuEdit	true	"body content"
// @Success 200 {string} models.Demo.Id
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router / [post]
func (d *DemoController) Post() {
	var demo models.Demo
	json.Unmarshal(d.Ctx.Input.RequestBody, &demo)
	valid := validation.Validation{}
	valid.Required(demo.Score, "分数").Message("不能为空")
	valid.Range(demo.Score, 0, 100, "分数").Message("需在0~100之间")
	valid.Required(demo.PlayerName, "选手姓名").Message("不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			d.CustomAbort(400, fmt.Sprintf("%s%s", err.Key, err.Message))
		}
	}
	id, err := models.AddDemo(demo)
	if err != nil {
		logs.Error(err.Error())
		d.CustomAbort(500, err.Error())
	}
	d.Data["json"] = map[string]string{"id": id}
	d.ServeJSON()
}

// @Title 获取某对象
// @Description 根据ID获取某对象文档
// @Param	id	path	string	true	"某对象ID"
// @Success 200 {object} models.DemoDocuEdit
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router /:id [get]
func (d *DemoController) Get() {
	id := d.GetString(":id")
	if bson.IsObjectIdHex(id) == false {
		d.CustomAbort(400, "某对象ID的格式不正确")
	}
	demo, err := models.GetDemo(id)
	if err != nil {
		logs.Error(err.Error())
		d.CustomAbort(500, err.Error())
	}
	d.Data["json"] = demo
	d.ServeJSON()
}

// @Title 获取全部某对象
// @Description 获取全部某对象文档
// @Success 200 {object} models.DemoDocuEdit
// @Failure 500 服务异常提示字符串
// @router / [get]
func (d *DemoController) GetAll() {
	demo, err := models.GetAllDemo()
	if err != nil {
		logs.Error(err.Error())
		d.CustomAbort(500, err.Error())
	}
	d.Data["json"] = demo
	d.ServeJSON()
}

// @Title 更新某对象
// @Description 根据ID更新某对象文档
// @Param	id	path	string	true	"某对象ID"
// @Param	body	body	models.DemoDocuUpdate	true	"body content"
// @Success 200 {string} Update Success!
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router /:id [put]
func (d *DemoController) Put() {
	id := d.GetString(":id")
	if bson.IsObjectIdHex(id) == false {
		d.CustomAbort(400, "某对象ID的格式不正确")
	}
	var object models.DemoDocuUpdate
	json.Unmarshal(d.Ctx.Input.RequestBody, &object)
	valid := validation.Validation{}
	valid.Required(object.Score, "分数").Message("不能为空")
	valid.Range(object.Score, 0, 100, "分数").Message("需在0~100之间")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			d.CustomAbort(400, fmt.Sprintf("%s%s", err.Key, err.Message))
		}
	}
	err := models.UpdateDemo(id, object)
	if err != nil {
		logs.Error(err.Error())
		d.CustomAbort(500, err.Error())
	}
	d.Data["json"] = "Update Success!"
	d.ServeJSON()
}

// @Title 删除某对象
// @Description 根据ID删除某对象文档
// @Param	id	path	string	true	"某对象ID"
// @Success 200 {string} Delete Success!
// @Failure 400 请求异常提示字符串
// @Failure 500 服务异常提示字符串
// @router /:id [delete]
func (d *DemoController) Delete() {
	id := d.GetString(":id")
	if bson.IsObjectIdHex(id) == false {
		d.CustomAbort(400, "某对象ID的格式不正确")
	}
	err := models.DeleteDemo(id)
	if err != nil {
		logs.Error(err.Error())
		d.CustomAbort(500, err.Error())
	}
	d.Data["json"] = "Delete Success!"
	d.ServeJSON()
}
