package main

import (
	"fmt"
	"ginDemo1/config"
	"ginDemo1/controller"
	"ginDemo1/datasource"
	"ginDemo1/service"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {

	app := newApp()

	configatino(app)

	//路由设置
	mvcHandle(app)

	config := config.InitConfit()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
	testfunc()

}

func newApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")

	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(ctx *context.Context) {
		ctx.View("index.html")

	})
	return app
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {

	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	engine := datasource.NewMysqlEngine()
	adminService := service.NewAdminService(engine)

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

}

/**
 * 项目配置
 */
func configatino(app *iris.Application) {
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	app.OnErrorCode(iris.StatusNotFound, func(ctx *context.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx *context.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}

func testfunc() {
	fmt.Println("hello golang")
}
