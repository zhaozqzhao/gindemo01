package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

func main() {

	app := iris.New()
	app.Get("/user", func(ctx *context.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		ctx.JSON(map[string]interface{}{
			"code": http.StatusOK,
			"msg":  "处理get请求成功，查询用户信息",
		})

	})

	mvc.New(app).Handle(new(UserController))

	app.Run(iris.Addr(":8080"))
}

type UserController struct{}

func (uc *UserController) GetInfo() mvc.Result {

	iris.New().Logger().Info("get请求，请求路径为info")
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 1,
			"msg":  "请求成功",
		},
	}

}
