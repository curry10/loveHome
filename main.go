package main

import (
	
	_ "loveHome/routers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	_ "loveHome/models"
)

func main() {
	//设置一个fastdfs 请求的静态路径
	//http://101.200.170.171:8080/group1/M00/00/00/Zciqq1oaGW-ABnxDAAAHFIcthTk%207176.go
	beego.SetStaticPath("/group1/M00", "fastdfs/storage_data/data")

	//测试fastdfs接口
	//	models.FDFSUploadByFileName("home01.jpg")

	ignoreStaticPath()
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.Run()
}

//重定向static静态路径
func ignoreStaticPath() {

	//透明static

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)

	//将全部的静态资源重定向 加上/static/html路径
	//http://ip:port:8080/index.html----> http://ip:port:8080/static/html/index.html
	//如果restFUL api  那么就取消冲定向
	//http://ip:port:8080/api/v1.0/areas ---> http://ip:port:8080/static/html/api/v1.0/areas
}