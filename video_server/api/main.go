package main

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	logs.Deb	ug("启动服务...")
	http.ListenAndServe(":8080", r)
}