//Golang学习手册之：带你21周搞定Go语言
//https://www.bilibili.com/video/BV16E411H7og?p=181
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin HelloWorld
func main(){
	//1. 创建路由
	r := gin.Default()
	//2. 绑定路由规则，执行的函数
	//gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "hello World!")
	})
	//3. 监听端口，默认8080
	r.Run(":8000")
}