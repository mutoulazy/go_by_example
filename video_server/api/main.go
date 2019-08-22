package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	logs.Debug("启动服务...")
	http.ListenAndServe(":8080", r)
}