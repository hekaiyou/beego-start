package controllers

import (
	"fmt"
	"time"
	"strings"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
)

// ErrorJSON 异常JSON数据类型
type ErrorJSON struct {
	Code int
	Message string
	ClientIP string
	ServerTime time.Time
}

/*
ErrorResponseJSON 异常响应JSON
	ctx		请求上下文类型
	code	请求状态码
	data	异常JSON数据
控制器设计
	errorJSON := ErrorJSON{
		Message: fmt.Sprintf("%s%s", err.Key, err.Message),
	}
	ErrorResponseJSON(c.Ctx, 400, errorJSON)
	return
*/
func ErrorResponseJSON(ctx *context.Context, code int, data ErrorJSON) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 判断请求是否来自于 Swagger 网页端
	if strings.Contains(ctx.Request.Header.Get("Referer"), "/swagger/") {
		ctx.ResponseWriter.WriteHeader(200)
	} else {
		ctx.ResponseWriter.WriteHeader(code)
	}
	data.Code = code
	data.ClientIP = ctx.Input.IP()
	// nginx 中 proxy_set_header 设置的值
	// data.ClientIP = ctx.Request.Header.Get("X-Real-ip")
	data.ServerTime = time.Now().UTC()
	ctx.Output.JSON(data, true, true)
}

/*
ValidationInspection 验证项检查方法
	ctx		请求上下文类型
	valid	验证对象
控制器设计
	valid := validation.Validation{}
	valid.Required(demo.Score, "得分").Message("不能为空")
	valid.Range(demo.Score, 0, 100, "得分").Message("需在0~100之间")
	valid.Required(demo.PlayerName, "选手姓名").Message("不能为空")
	if ValidationInspection(c.Ctx, valid) != nil {
		return
	}
*/
func ValidationInspection(ctx *context.Context, valid validation.Validation) error {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorJSON := ErrorJSON{
				Message: fmt.Sprintf("%s%s", err.Key, err.Message),
			}
			ErrorResponseJSON(ctx, 400, errorJSON)
			return err
		}
	}
	return nil
}