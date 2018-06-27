package main

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"runtime"
	"zhulei.com/stone/context"
	"zhulei.com/stone/db"
	"zhulei.com/stone/routes/api"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer db.Pool().Close()

	e := echo.New()
	e.SetHTTPErrorHandler(context.HTTPErrorHandler)
	configRoutes(e)

	server := standard.New(":8090")
	//e.Run(server)
	server.SetHandler(e)
	server.SetLogger(e.Logger())
	gracehttp.Serve(server.Server)
}

func configRoutes(e *echo.Echo) {
	api := e.Group("/api")

	//注册
	api.POST("/register", stone.Register)
}
